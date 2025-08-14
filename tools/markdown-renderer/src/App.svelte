<script>
    import MarkdownRender from "./lib/MarkdownRender.svelte";
    import Textbox from "./lib/Textbox.svelte";
    let md = $state();
</script>

<main class="ml-6 mr-6">
    <div class="content ml-3">
        <h1 class="is-size-1 mt-3">Markdown Renderer</h1>
        <h3>Just paste some markdown and watch it transform in real time!</h3>
    </div>
    <div class="is-flex is-flex-direction-row is-justify-content-space-between">
        <div class="left-side">
            <Textbox bind:markdown={md}></Textbox>
        </div>
        <div class="right-side">
            <MarkdownRender markdown={md}></MarkdownRender>
        </div>
    </div>
</main>

<style>
    .left-side {
        width: 45%;
    }
    .right-side {
        width: 55%;
    }
</style>
