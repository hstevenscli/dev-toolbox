<script>
    import showdown from "showdown";
    import "github-markdown-css/github-markdown.css"

    let { markdown } = $props();
    let converter = new showdown.Converter()
    let html = $derived(converter.makeHtml(markdown));

</script>
<link rel="stylesheet" href="github-markdown.css">

<article class="markdown-body">
    <div>{@html html}</div>
</article>

<style>
	.markdown-body {
		box-sizing: border-box;
		min-width: 200px;
		max-width: 980px;
		margin: 0 auto;
        padding: 15px;
	}

	@media (max-width: 767px) {
		.markdown-body {
            padding: 5px;
		}
	}
    article :global {
        ul {
            list-style: disc;
        }
    }
</style>
