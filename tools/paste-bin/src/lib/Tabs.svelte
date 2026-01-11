<script>

    let { tabs } = $props()
    let currentTab = $state(0);

    function selectTab(index) {
        currentTab = index;
    }
</script>

<div class="container">
{#each tabs as tab, i }
    <div onclick="{() => {
        selectTab(i);
        }}" class="tab { i == currentTab ? 'is-underline' : ''}">
        {tab} 
    </div>
    {#if i != tabs.length - 1} &nbsp;&nbsp;<div class="vbar"></div>&nbsp;&nbsp;{/if}
{/each}
</div>

<style>
    .vbar {
        width: 3px;        /* thickness */
        height: 43px;
        background: grey;  /* color */
    }
    .tab {
        background-color: none;
        padding: 8px;
        border: 1px solid grey;
        border-radius: 3px;
    }
    .is-underline {
        border-bottom: 3px solid white;
    }
    .tab:hover {
        background-color: grey;
        outline: none;
        cursor: pointer;
    }
    .tab:active {
        background-color: lightblue;
    }
    .tab:focus {
        outline: none;
    }
    .container {
        display: flex;
        flex-direction: row;
        align-items: center;
    }
</style>
