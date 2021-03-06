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

	"github.com/anacrolix/torrent"
	"github.com/spf13/cobra"
)

var filename string
var dirdata string

// torrentCmd represents the torrent command
var torrentCmd = &cobra.Command{
	Use:   "torrent",
	Short: "Download data from a .torrent file",
	Long:  `Download data from a .torrent file in a secure way`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := torrent.NewDefaultClientConfig()
		cfg.DataDir = dirdata
		cfg.Debug = false
		cfg.NoUpload = true

		c, err := torrent.NewClient(cfg)
		if err != nil {
			log.Fatalln("[X]Error creating Torrent Client")
		}
		defer c.Close()

		trrnt, err := c.AddTorrentFromFile(filename)
		if err != nil {
			log.Fatalln("[X]Error reciving the torrent")
		}

		<-trrnt.GotInfo()
		trrnt.DownloadAll()

		zipfile := trrnt.Name()

		c.WaitAll()
		log.Printf("File %s is completed download", zipfile)
		log.Println("[!]Please Wait Wile Clean Your Folder")

		pathtoremove := datadir + ".torrent.*"
		defer os.Remove(pathtoremove)
	},
}

func init() {
	rootCmd.AddCommand(torrentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// torrentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// torrentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	torrentCmd.Flags().StringVarP(&filename, "file", "f", "", "use this flag for set the torrent file to download")
	torrentCmd.Flags().StringVarP(&dirdata, "data", "d", "./", "use this flag for set the download folder")
}
