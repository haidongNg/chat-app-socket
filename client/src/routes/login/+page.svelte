<script>
    import "../../app.css";
</script>

<div class="h-full w-full flex justify-center items-center bg-indigo-300">
    <div
        class="bg-white shadow-lg py-5 px-10 border rounded-lg flex flex-col gap-2 items-center"
    >
        <span class="font-bold text-2xl text-indigo-700">QuickChat</span>
        <span class="text-xs text-indigo-300">Login</span>
        <form class="flex flex-col gap-4">
            <input type="email" placeholder="Email" />
            <input type="password" placeholder="Password" />
            <button
                class="bg-indigo-400 hover:bg-indigo-500 text-white font-bold p-2 border-none cursor-pointer rounded-lg"
                >Sign up</button
            >
        </form>
        <p>ou don't have an account? <a href="/register">Register</a></p>
    </div>
</div>