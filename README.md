# GOTORRENT
is a tool for download magnet and torrent files in a safe way

## INSTALL

+ `git clone https://github.com/Valkirian/r4st4t0rr3nt.git && cd r4st4t0rr3nt`
+ `go build`

## USAGE GENERAL
```
Usage:
  GoTorrent [command]

Available Commands:
  help        Help about any command
  magnet      command for download a torrent for magnet url
  torrent     A brief description of your command

Flags:
  -h, --help     help for GoTorrent
  -t, --toggle   Help message for toggle

Use "GoTorrent [command] --help" for more information about a command.
```

### USAGE WITH MAGNET LINK
```
command for download a torrent for magnet url using a --url flag for define the file to download

Usage:
  GoTorrent magnet [flags]

Flags:
  -d, --data string   set this flag for storage your data in other folder (default "./")
  -h, --help          help for magnet
  -u, --url string    set the magnet url to download file
```

### USAGE WITH .TORRENT FILE
```
Download data from a .torrent file in a secure way

Usage:
  GoTorrent torrent [flags]

Flags:
  -d, --data string   use this flag for set the download folder (default "./")
  -f, --file string   use this flag for set the torrent file to download      
  -h, --help          help for torrent
```
