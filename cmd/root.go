package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	//"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var Throttling int64

var rootCmd = &cobra.Command{
	Use:   "hrf",
	Short: "Hash remote file is a way of obtaining the hash of a file without storing it",
	//Need better wording
	Long: `A way to target a remote file and obtain its hash without storing it into you compute.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Pass an URL after the options (e.g: http//www.google.com")
		}
		_, err := url.ParseRequestURI(args[0])
		if err != nil {
			log.Print(err)
			log.Fatalf("Pass a correft URL after the options (e.g: http//www.google.com )")
		}
		resp, err := http.Get(args[0])
		if err != nil {
			log.Fatalf("Get failed: %v", err)
		}
		defer resp.Body.Close()
		//wrappedIn := flowrate.NewReader(resp.Body, Throttling)

		chanWriter := NewChanWriter()
		// Copy to channel
		go func() {
			defer chanWriter.Close()
			chanWriter.Write([]byte{12})
			//_, err = io.Copy(chanWriter, wrappedIn)
		}()


		var hash []byte
		for c := range chanWriter.Chan() {
			hash = Hash(c)
			fmt.Println(hash)
		}

		if err != nil {
			log.Fatalf("Copy failed: %v", err)
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.Flags().Int64VarP(&Throttling, "throttling", "t", 0, "Throttle the download to a rate of bytes. Pass a number that will be parsed as the number of KB that this will be throttled")
}