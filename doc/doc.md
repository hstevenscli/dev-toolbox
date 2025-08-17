# Toolbox Server

Note: "/" represents the root of the project directory

Backend located in 
```
/
```

Main landing page
```
/client
```


Each tool should be in its own directory and then hooked up to the server or main client for production

```
ProjectRoot
├──main.go
├──handlers.go
├──go.mod
├──go.sum
├──db
│   ├── db1
│   ├── db2
│   └── db3
├──client
│   ├── index.html/js
│   └── style.css
└──tools  ← SEPARATE TOOLBOX TOOLS GO HERE
    ├── json formatter
    ├── metadata editor
    ├── pdf tools (img to pdf, md to pdf, etc)
    └── markdown renderer
```


## Tool Location and Integration

Each tool should be in its own folder inside of the 'tools' directory. Each tool should consist of some frontend files (index.html, style.css, \*.js)

If backend logic is needed for the tool, a .go file of the same name should be created in the root directory.

Example:

If I have mdRenderer/ inside of the tools directory I would create mdRenderer.go at the root directory and have any backend logic there.

## Do I have to use svelte?

No, I used svelte for the client directory and the tools that I have built, but you should be able to use any frontend technology that uses npm to run (react, vue, svelte, preact, etc.)

Other frontend frameworks shopuld work as well but they might need some manual tweaking to get them to mesh with the rest of the project.

You could also skip frameworks and build steps altogether and just have a html/css/js project as well


## How to ensure tools are included in the build

Im trying to automate all of this but as of now some links need to be manually made in key locations

1. Each tool needs to have a package.json or similar file in its directory

2. A link also needs to be included in ./client/src/App.svelte

Example:
```html
    <a href="/cooltool">Cool Tool</a>
```
3. A matching handler function will need to be included in the main.go file to serve that route as well

Example:

```go
	router.GET("/cooltool", getCoolToolPage)
```

This should ensure that the tool gets built and included as part of the larger project on deployment


## DB

Don't know how to best handle data persistence yet. For now we have a db directory that can be used however we happen to do it
