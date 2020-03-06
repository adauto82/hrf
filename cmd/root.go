package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/mxk/go-flowrate/flowrate"
	"github.com/spf13/cobra"
	"io"
	"path/filepath"

	"log"
	"net/http"
	"net/url"
	"os"
)

const DEFAULT_FILE_NAME = "default_file"

var Throttling int64
var FilePath string

var rootCmd = &cobra.Command{
	Use:   "hrf [OPTIONS] URL",
	Short: "Hash remote file is a way of obtaining the hash of a file without storing it",
	//Need better wording
	Long: `A way to target a remote file and obtain its hash without storing it into you compute. 
	If no --file_path is passed the hash will be stored in the default_file on your CWD.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Pass an URL after the options (e.g: http//www.google.com )")
		}
		_, err := url.ParseRequestURI(args[0])
		if err != nil {
			log.Print(err)
			log.Fatalf("Pass a correct URL after the options (e.g: http//www.google.com )")
		}

		inf, err := os.Stat(FilePath)
		//either is a file or it is a directory
		if err == nil {
			switch mode := inf.Mode(); {
			case mode.IsDir():
				FilePath = filepath.Join(FilePath, DEFAULT_FILE_NAME)
			}
		}
		f, err := os.Create(FilePath)
		if err != nil {
			log.Fatalf("Could not create the file in the required path")
		}

		defer f.Close()

		resp, err := http.Get(args[0])
		if err != nil {
			log.Fatalf("Get failed: %v", err)
		}
		defer resp.Body.Close()
		wrappedIn := flowrate.NewReader(resp.Body, Throttling)

		chanWriter := NewChanWriter()
		// Copy to channel
		go func() {
			defer chanWriter.Close()
			_, err = io.Copy(chanWriter, wrappedIn)
			if err != nil {
				log.Fatalf("Copy failed: %v", err)
			}
		}()

		var hash []byte
		for c := range chanWriter.Chan() {
			hash = Hash(c)
		}
		hexString := hex.EncodeToString(hash)
		//fmt.Println(hexString)
		_, err = f.WriteString(hexString)
		if err != nil {
			log.Fatalf("Error writing HEX (%s) to file", hexString)
		}
		f.Sync()
		log.Println("File " + f.Name() + " writed with the hex of the URL downloaded")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.Flags().StringVarP(&FilePath, "file_path", "p", DEFAULT_FILE_NAME, "The path to where the hash will be written")
	//TODO: Permit KB / MB
	rootCmd.Flags().Int64VarP(&Throttling, "throttling", "t", 0, "Throttle the download to a rate of bytes. If no number is passed, then no limits.")
}
