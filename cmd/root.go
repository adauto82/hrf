package cmd

import (
	"fmt"
	"github.com/mxk/go-flowrate/flowrate"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hrf",
	Short: "Hash remote file is a way of obtaining the hash of a file without storing it",
	//Need better wording
	Long: `A way to target a remote file and obtain its hash without storing it into you compute.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://google.com")
		if err != nil {
			log.Fatalf("Get failed: %v", err)
		}
		defer resp.Body.Close()
		wrappedIn := flowrate.NewReader(resp.Body, 10)

		// Copy to stdout
		_, err = io.Copy(os.Stdout, wrappedIn)
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
	rootCmd.PersistentFlags().Bool("throttling", true, "use download throttling")
}