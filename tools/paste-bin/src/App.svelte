<script>
    import Tabs from "./lib/Tabs.svelte";


    async function getPastebin() {
        let response = await fetch("/db/pastebin");
        let json = await response.json();
        entries = json;
    }
    let entries = $state("");
    let tabslist = ["PASTE", "RECORDS", "SETTINGS"]
</script>

<Tabs tabs={tabslist}/>


<p>Helloworld</p>
<div class="text-and-buttons">
    <textarea></textarea>
    <div class="buttons">
        <button class="button submit">Submit</button>
        <button class="button clear">Clear</button>
    </div>
    <p class="entries-list">{ entries }</p>
    <button onclick={getPastebin} class="button">Get Pastebin</button>
</div>

<style>
    .submit {
        background-color: green;
    }
    .clear {
        background-color: red;
    }
    .buttons {
        display: flex;
        flex-direction: row;
        justify-content: center;
    }
    .button {
        margin: 3px;
    }
    .text-and-buttons {
        width: auto;
    }
    .button:hover {
        background-color: grey;
        border: none;
    }
    .button:active {
        background-color: lightblue;
    }
</style>
