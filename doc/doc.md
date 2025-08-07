# Toolbox Server

Note: "/" represents the root of the project directory

Backend located in 
```
/server
```

Main landing page
```
/client
```


Each tool should be in its own directory and then hooked up to the server or main client for production

```
ProjectRoot
├──db
│   ├── db1
│   ├── db2
│   └── db3
├──server
│   ├── main.go
│   ├── handlers.go
│   └── etc
├──client
│   ├── index.html/js
│   └── style.css
└──tools  ← Separate toolbox tools go here
    ├── json formatter
    ├── metadata editor
    ├── pdf tools (img to pdf, md to pdf, etc)
    └── markdown renderer
```


## Tools

Each tool should be in its own folder inside of the 'tools' directory. Each tool should consist of some frontend files (index.html, style.css, \*.js)

If backend logic is needed for the tool, a .go file of the same name should be created in the 'server' directory.

Example:

If I have mdRenderer/ inside of the tools directory I would create mdRenderer.go inside of the server directory and have any backend logic in that folder.

## DB

Don't know how to best handle data persistence yet. For now we have a db directory that can be used however we happen to do it
