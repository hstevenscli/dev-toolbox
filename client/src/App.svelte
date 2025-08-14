<script>
    let links = ["MDRender", "Json Formatter"]
</script>

<main>
    <h1>Welcome to the ToolBox!</h1>
    <br>
    <br>

    <h3>Helpful Tools</h3>

    {#each links as link}
        &gt; <a title="Some info here" href="/tools/{link}">{link}</a><br>
    {/each}
    <a href="/markdownrender">MD</a>
</main>

<style>
    ul {
        list-style: none;
    }
    button {
        background-color: green;
    }
</style>
