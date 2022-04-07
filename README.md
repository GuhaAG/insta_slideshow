# Instagram Slideshow
Go API to fetch instragram media with a username, and display them on a webpage. 

### dependencies
- yarn
- npm
- go
- dep

### Quickstart 
```bash
go get github.com/guhaag/insta_slideshow
cd $GOPATH/src/github.com/guhaag/insta_slideshow

build:
	go build
build_fe:
	cd client && yarn && yarn build	
run:
	go run App.go
```

# Project Structure
```bash
.
├── App.go
├── Gopkg.toml
├── README.md
├── backend
│   └── Backend.go
├── client
│   ├── package.json
│   ├── public
│   │   ├── favicon.ico
│   │   ├── index.html
│   │   └── manifest.json
│   ├── src
│   │   ├── App.css
│   │   ├── App.js
│   │   ├── App.test.js
│   │   ├── components
│   │   │   └── demo
│   │   │       └── demo.js
│   │   ├── index.css
│   │   ├── index.js
│   │   └── logo.svg
│   └── yarn.lock
└── frontend -> client/build
```

The server is running at [localhost:3000](http://localhost:3000/). You can see the API at [localhost:3000/api/v1/users](http://localhost:3000/api/v1/users)
