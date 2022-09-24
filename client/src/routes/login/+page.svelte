<script>
    import { user } from "../../stores/user";
    import { goto } from "$app/navigation";
    
    let dataForm = {
        email: "",
        password: "",
    };
    
    /**
     * Login
     */
    const handleLogin = () => {
        user.set(dataForm);
        goto('/');
    };
</script>

<svelte:head>
    <title>Login</title>
</svelte:head>
<div class="h-full w-full flex justify-center items-center bg-indigo-300">
    <div
        class="bg-white shadow-lg py-5 px-10 border rounded-lg flex flex-col gap-2 items-center"
    >
        <span class="font-bold text-2xl text-indigo-700">QuickChat</span>
        <span class="text-xs text-indigo-300">Login</span>
        <form
            class="flex flex-col gap-4"
            on:submit|preventDefault={handleLogin}
        >
            <input
                type="email"
                placeholder="Email"
                bind:value={dataForm.email}
            />
            <input
                type="password"
                placeholder="Password"
                bind:value={dataForm.password}
            />
            <button
                class="bg-indigo-400 hover:bg-indigo-500 text-white font-bold p-2 border-none cursor-pointer rounded-lg"
                >Sign up</button
            >
        </form>
        <p>You don't have an account? <a href="/register">Register</a></p>
    </div>
</div>