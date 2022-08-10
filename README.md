# go-websites
This repo is a constant work in progress to provide pointers for creating web applications in Go.

## Starting Point
If you are creating multiple websites with Go it's good to create a uniformed folder structure. I took some pointers from [@AlexEdwards](https://github.com/alexedwards) when I went through his [Let's Go](https://lets-go.alexedwards.net/) Course. I highly recommend his course when learning Go for web applications. 

This repo contains two scripts that can be used to quickly create that folder structure for you.

You can execute either one based on preference:

    Go file: setup-go-website.go
    PowerShell file: setup-go-website.ps1

Both files take a [WebsiteName] parameter to create a base folder and then create the subfolders inside of that.

### Go File
```go
go run setup-go-website.go -websiteName <website>
```
If a website name is not defined a default folder will be created for you: ```default-go-website```

### PowerShell File
```powershell
PS /> .\setup-go-website.ps1 -websiteName <website>
```

Once you execute one of these files you will be left with a folder structure like this:
```
website
│
└───cmd
│   │
│   └───web
│       │   handlers.go
│       │   main.go
│   
└───internal
│
└───ui
    └───html
    │   └───pages
    │   └───partials
    │ 
    └───static
        └───css
        └───img
        └───js
```

## Structure Defined
Your website name will be the root folder.

The next folder inside will be cmd/web which will house your base go files: main.go and handlers.go. 

The internal folder can hold potentially re-usable code or items that should not be exposed. 

The ui folder per definition is user interface. Any html files and static files will be under this portion of the tree. 

The html folder will house pages that can define specific page templates. Partials can house specific page elements that you want to be repeated such as navigation. 

The static folder will contain items such as css, any images, and javascript files (js). 