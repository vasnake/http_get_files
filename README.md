# http_get_files

Very simple Go program, enumerate files in given directory and serve them via HTTP GET

Motivation: I tried to setup VLC streaming (server) on my win11 box,
just to have an ability to watch a few TV series (files on my win11 box) on my phone.
Well, this time VLC magic dindn't work for me.
After a couple of hours I said to myself - "fuck it, all I need is a primitive HTTP server, nothing more".
And there it is, a primitive HTTP server.

I build it and run in WSL2 under Win11 env.

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

# network, Win11 + WSL2

# ubuntu
ifconfig # 172.19.24.89

# win
netsh advfirewall firewall add rule name="Allowing LAN 8080" dir=in action=allow protocol=TCP localport=8080
netsh interface portproxy add v4tov4 listenaddress=0.0.0.0 listenport=8080 connectaddress=172.19.24.89 connectport=8080
netsh interface portproxy show v4tov4
netsh interface ip show address

```
snippets
