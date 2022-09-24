<script lang="ts">
  import Sidebar from "../components/Sidebar.svelte";
  import Chat from "../components/Chat.svelte";
</script>

<svelte:head>
  <title>Home</title>
</svelte:head>
<div class="h-full w-full flex">
  <div class="flex overflow-hidden w-full">
    <Sidebar />
    <Chat />
  </div>
</div>
