/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"os"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var url string
var datadir string

// magnetCmd represents the magnet command
var magnetCmd = &cobra.Command{
	Use:   "magnet",
	Short: "command for download a torrent for magnet url",
	Long:  `command for download a torrent for magnet url using a --url flag for define the file to download`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := torrent.NewDefaultClientConfig()
		cfg.DataDir = datadir
		cfg.Debug = false
		cfg.NoUpload = true

		c, err := torrent.NewClient(cfg)
		if err != nil {
			log.Fatalln("[X]Error creating Torrent Client")
		}
		defer c.Close()

		trrnt, err := c.AddMagnet(url)
		if err != nil {
			log.Fatalln("[X]Error reciving the torrent")
		}
		<-trrnt.GotInfo()
		trrnt.DownloadAll()

		zipfile := trrnt.Name()
		size := trrnt.Length()

		//Progress bar
		bar := progressbar.DefaultBytes(size)
		for i := 0; i < int(size); {
			bar.Add64(trrnt.BytesCompleted())
			time.Sleep(1000 * time.Millisecond)
		}

		c.WaitAll()
		log.Printf("File %s is completed download", zipfile)
		log.Println("[!]Please Wait Wile Clean Your Folder")

		pathtoremove := datadir + ".torrent.*"
		defer os.Remove(pathtoremove)
	},
}

func init() {
	rootCmd.AddCommand(magnetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// magnetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// magnetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	magnetCmd.Flags().StringVarP(&url, "url", "u", "", "set the magnet url to download file")
	magnetCmd.Flags().StringVarP(&datadir, "data", "d", "./", "set this flag for storage your data in other folder")
}
