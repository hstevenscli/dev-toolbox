<script>
    let { message = 'Paste in the ol\' markdown here', markdown = $bindable() } = $props();

    let rows = $state(5)
    let textAreaEl;
    //let lines = $state(textAreaEl.value)

    //let rows = $state(5)
    //function resizeTextArea() {
    //    let lines = markdown.split('\n');
    //    rows = Math.max(5, lines.length);
    //}

    function resize() {
        rows = Math.max(5, textAreaEl.value.split('\n').length);
    }


</script>

<textarea rows="{rows}" bind:this={textAreaEl} bind:value={markdown} class="textarea" placeholder="{message}" oninput={resize}></textarea>

<style>
    textarea {
        height: 10000px;
    }
</style>
