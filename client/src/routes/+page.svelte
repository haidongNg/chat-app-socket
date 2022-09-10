<script lang="ts">
	import { ENDPOINT } from '$lib/realtime';
	import { onMount } from 'svelte';
	import '../app.css';

	let textField = '';
	let username = Date.now().toString(36) + Math.random().toString(36).substr(2);
	let messages: any[] = [];
	let conn: WebSocket;
	onMount(() => {
		console.log('the component has mounted');
		if (window['WebSocket']) {
			conn = new WebSocket(ENDPOINT + 'global'+'?roomId=global&roomName=global'+`&clientId=${username}`);
			conn.onmessage = (evt) => {
				console.log(evt);
				const data = JSON.parse(evt.data);
				messages = [...messages, data];
			};
		}
	});

	// Send
	function sendMessage() {
		if (!conn) return;

		const message = textField.trim() + username;

		if (!message) return;

		textField = '';
		// Send the message
		conn.send(
			JSON.stringify({
				message: message,
				clientId: username,
				roomId: 'global'
			})
		);
	}
</script>

<div class="bg-[#F2F5F8] p-12 h-screen w-screen text-[#434651]">
	<a href="/room">Room my site</a>
	<input type="text" placeholder="Username" bind:value={username} />
	<div class="p-8 max-h-96 overflow-auto">
		{#each messages as message}
			{#if message.clientId === username}
			<div class="bg-[#94C2ED] border rounded-md p-2 m-2 after:content-[' '] after:clear-both after:table">
				<img
					class="h-12 w-12 rounded-full float-right  mr-0 ml-5"
					src="https://images.unsplash.com/photo-1491528323818-fdd1faba62cc?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
					alt=""
				/>
				<p>{message.message}</p>
				<span class="float-left">11:00</span>
			</div>
			{:else} 
			<div class="bg-[#86BB71] border rounded-md p-2 m-2 after:content-[' '] after:clear-both after:table">
				<img
					class="h-12 w-12 rounded-full float-left mr-5"
					src="https://images.unsplash.com/photo-1491528323818-fdd1faba62cc?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
					alt=""
				/>
				<p>{message.message}</p>
				<span class="float-right">{new Date().toDateString()}</span>
			</div>
			{/if}
			
		{/each}
		
	</div>

	<form on:submit|preventDefault={sendMessage} class="p-8">
		<textarea
			class="w-full border-none outline-none py-2 px-5 mb-2 rounded-md resize-none"
			name=""
			id=""
			cols="30"
			rows="10"
			placeholder="Type your message"
			bind:value={textField}
		/>
		<button
			class="text-[#94C2ED] bg-[#F2F5F8] hover:text-[#94C2ED] float-right uppercase border-none cursor-pointer font-bold "
			>Send</button
		>
	</form>
</div>
