# http_get_files

Very simple Go program, enumerate files in given directory and serve them via HTTP GET

Motivation: I tried to setup VLC streaming (server) on my win11 box,
just to have an ability to watch a few TV series (files on my win11 box) on my phone.
Well, this time VLC magic dindn't work for me.
After a couple of hours I said to myself - "fuck it, all I need is a primitive HTTP server, nothing more".
And there it is, a primitive HTTP server.

On my phone, in media player I just open HTTP stream from URL like `http://192.168.1.6:8080/1` and voilà,
I'm watching first episode from selected season (directory) of my favorite TV show.

Building and running in WSL2 under Win11 env:

Build
```s
pushd /mnt/c/Users/vlk/data/github/http_get_files/
PATH=${HOME}/go/bin:${PATH}

go mod init http_get_files

# makes go vet happy
cat > main.go << EOT
package main
func main() { panic("not yet") }
EOT

go mod tidy
gofmt -w .
go vet http_get_files
go test http_get_files
go run http_get_files

# network setup, Win11 + WSL2

# ubuntu
ifconfig # 172.19.24.89

# win
netsh advfirewall firewall add rule name="Allowing LAN 8080" dir=in action=allow protocol=TCP localport=8080
netsh interface portproxy add v4tov4 listenaddress=0.0.0.0 listenport=8080 connectaddress=172.19.24.89 connectport=8080
netsh interface portproxy show v4tov4
netsh interface ip show address

```
snippets

Program have two parameters: HTTP server port, and path to the directory with files to serve.
Usage example: `go run http_get_files -port 8080 -path "/mnt/c/Users/vlk/Downloads/TV_series"`

All files in given directory will be sorted, enumerated starting from `1`,
and made accessable via urls like `http://192.168.1.6:8080/1`
where `192.168.1.6` is IP address for your box, and `8080` is IP port given in parameters.

Output example
```s
# wsl2
go run http_get_files -port 8080 -path /mnt/c/Users/vlk/Downloads/TV
2024-08-28T11:16:46.724Z: Parameters: port, dir: "8080"; "/mnt/c/Users/vlk/Downloads/TV"; 
2024-08-28T11:16:46.726Z: skip directory "Eng";
2024-08-28T11:16:46.726Z: file added: "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E01.avi";
2024-08-28T11:16:46.726Z: file added: "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E02.avi";
2024-08-28T11:16:46.726Z: file added: "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E03.avi";
2024-08-28T11:16:46.726Z: file added: "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E04.avi";
2024-08-28T11:16:46.726Z: file added: "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E05.avi";
2024-08-28T11:16:46.726Z: file added: "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E06.avi";
2024-08-28T11:16:46.726Z: file added: "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E07.avi";
2024-08-28T11:16:46.726Z: file added: "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E08.avi";
2024-08-28T11:16:46.726Z: added mapping "1"; "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E01.avi";
2024-08-28T11:16:46.726Z: added mapping "2"; "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E02.avi";
2024-08-28T11:16:46.726Z: added mapping "3"; "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E03.avi";
2024-08-28T11:16:46.726Z: added mapping "4"; "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E04.avi";
2024-08-28T11:16:46.726Z: added mapping "5"; "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E05.avi";
2024-08-28T11:16:46.726Z: added mapping "6"; "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E06.avi";
2024-08-28T11:16:46.726Z: added mapping "7"; "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E07.avi";
2024-08-28T11:16:46.726Z: added mapping "8"; "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E08.avi";
2024-08-28T11:16:46.726Z: Files mapping for server, Json:
{
  "1": "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E01.avi",
  "2": "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E02.avi",
  "3": "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E03.avi",
  "4": "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E04.avi",
  "5": "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E05.avi",
  "6": "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E06.avi",
  "7": "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E07.avi",
  "8": "/mnt/c/Users/vlk/Downloads/TV/foobar-S1E08.avi"
}
2024-08-28T11:16:46.726Z: Starting server at ":8080";
```
snippets
