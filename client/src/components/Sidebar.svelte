<script>
    import ChatList from "./ChatList.svelte";
    import Navbar from "./Navbar.svelte";
    // import Search from "./Search.svelte";
</script>

<div class="flex-1 bg-indigo-400  min-w-[270px] max-w-[270px]">
    <div class="flex flex-col py-8 px-6 bg-white flex-shrink-0 h-full">
        <Navbar />
        <!-- <Search /> -->
        <ChatList />
    </div>
</div>
