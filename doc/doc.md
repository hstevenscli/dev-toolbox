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

2. Make sure that assets are served at /tools/toolname/assets/assetname

You can do this manually or by changing your vite.config.js to have the following line:
> base: "/tools/tool-name/"

Example with markdown renderer
```javascript
import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vite.dev/config/
export default defineConfig({
    plugins: [svelte()],
    base: "/tools/markdown-renderer/" // added this line to the default config file
})

```

3. A link also needs to be included in ./client/src/App.svelte

Example:
```html
    <a href="/cooltool">Cool Tool</a>
```
4. The paths should be handled automatically in the go code, however you can manually set the path as a redundancy.

Include the directory name of your tool in the array in initLinks in handlers.go

Example:

```go
    // Calls getLinksLocal to get the names of all the tools in tools/
    func initLinks() []string {
        var links []string
        if l, err := getLinksLocal(); err != nil {
            links = []string{ "markdown-renderer", "json-formatter" } // <<<<<<< add the directory name here
        } else {
            links = l
        }
        return links
    }
```

You could also make a router.GET and handler func, but so long as you have dist/index.html in your tool directory it should be found and served automatically

If your project uses routes other than GET you will need to set those up manually for now

## DB

Don't know how to best handle data persistence yet. For now we have a db directory that can be used however we happen to do it
