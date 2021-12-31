# Go server on fiber

[fiber](https://github.com/gofiber/fiber)

## Install go

```sh
wget https://go.dev/dl/go1.17.5.linux-amd64.tar.gz  (linux)
wget https://go.dev/dl/go1.17.5.linux-arm64.tar.gz  (raspbery pi)
```

```sh
sudo tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz
```

Add to $HOME/.profile
```sh
sudo nano $HOME/.profile
```
```
export  PATH=$PATH:/usr/local/go/bin
```

```sh
source ~/.profile
```

## Build options

`go tool dist list` to list possible platforms

`go env GOOS GOARCH` to see env vars

Linux:
```sh
GOOS=linux GOARCH=amd64 go build
```

Windows:
```sh
GOOS=windows GOARCH=amd64 go build 
```

## VPS copy

```sh
scp jobs  root@example.com:/var/www/go
```

```sh
sudo nano /etc/systemd/system/jobs.service
```

Add this lines:
```
Description= instance to serve jobs api
After=network.target

[Service]
User=root
Group=www-data
ExecStart=/home/path/to/binary/you/uploaded(which in my case is/var/www/go/jobs)

[Install]
WantedBy=multi-user.target
```

```sh
sudo systemctl start jobs  
sudo systemctl enable jobs
```

## Nginx config

```sh
sudo nano /etc/nginx/sites-available/jobs
```

```
server {
listen 80 default_server;
listen [::]:80 default_server;
server_name example.com www.example.com ;  
location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

```sh
sudo ln -s /etc/nginx/sites-available/jobs /etc/nginx/sites-enabled
```

```sh
sudo nginx -t
sudo systemctl restart nginx
```