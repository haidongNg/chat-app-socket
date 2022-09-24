<script lang="ts">
    import { onMount } from "svelte";
    import Message from "./Message.svelte";
    let conn: WebSocket;

    onMount(() => {
        
    });
</script>

<div class="bg-gray-100 p-4 h-[calc(100%-112px)] overflow-y-scroll overflow-x-hidden">
    <Message owner={true} />
    <Message />
    <Message owner={true} />
    <Message />
    <Message owner={true} />
    <Message />
    <Message owner={true} />
    <Message />
    <Message />
    <Message owner={true} />
    <Message />
</div>
