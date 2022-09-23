<script>
  import Sidebar from "../components/Sidebar.svelte";
  import "../app.css";
  import Chat from "../components/Chat.svelte";
</script>

<div class="h-full w-full flex">
  <!-- container -->
  <div class="flex overflow-hidden w-full">
    <Sidebar />
    <Chat />
  </div>
</div>
 