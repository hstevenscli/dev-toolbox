<script>
    // let links = fetch("/links").then(response => response.json())
    async function getLinks() {
        let response = await fetch("/links");
        let json = await response.json();
        return json;
    }

    let links = getLinks();
</script>

<main>
    <h1>Welcome to the ToolBox!</h1>
    <br>
    <br>
    <h3>Helpful Tools</h3>

    {#await links}
        <!-- promise is pending -->
        <p>Getting Links...</p>
    {:then value}
        <!-- promise was fulfilled or not a Promise -->
        {#each value as link}
            &gt; <a title="Some info here" href="/{link}">{link}</a><br>
        {/each}
    {:catch error}
        <!-- promise was rejected -->
        <p>Something went wrong: {error.message}</p>
    {/await}
</main>

<style>
    ul {
        list-style: none;
    }
    button {
        background-color: green;
    }
</style>
