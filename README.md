# ipecho
A dumb server to tell you your IP on your network, or your IP as the world sees it

## Installation

These are the steps I use for raspbian. Some paths may differ slightly on other distros.

1. Build the executables
```
$ go build
```

2. Copy the executable to /bin
```
# cp ipecho /bin
```

3. Copy the .service file to systemd's config directory
```
# cp ipecho.service /etc/systemd/user
```

4. Reload the systemd daemon
```
$ systemctl --user daemon-reload
```

5. Start the service
```
# systemctl start ipecho.service
```

## Usage
By default, the service runs on port 5353. This can be changed in the service file via the `-port` flag.

To get your IP on your LAN or whatever, just visit the server.
```
$ curl ipecho.your.net:5353
```

To spice things up and get your external IP as the world (or at least AWS (which is the same thing)) sees it:
```
$ curl ipecho.your.net:5353/external
```
This endpoint just hits [checkip.amazonaws.com](http://checkip.amazonaws.com)


## Motivation
This solves a problem I have approximately every 8 to 14 months and is written in Go. It would probably be useful in a script.