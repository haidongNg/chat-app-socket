<script lang="ts">
	import { ENDPOINT } from '$lib/realtime';
	import { onMount } from 'svelte';
	import '../../app.css';

	let textField = '';
	let username = '';
	let messages: any[] = [];
	let conn: WebSocket;
	onMount(() => {
		console.log('the component has mounted');
		if (window['WebSocket']) {
			conn = new WebSocket(ENDPOINT + 'global/'+'123');
			console.log(messages);
			conn.onmessage = (evt) => {
				console.log(evt)
				const data = { message: evt.data, from: 'asdasd', time: '123123' };
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
		conn.send(message);
	}
</script>

<div class="bg-[#F2F5F8] p-12 h-screen w-screen text-[#434651]">
	<input type="text" placeholder="Username" bind:value="{username}">
	<div class="p-8 max-h-96 overflow-auto">
		{#each messages as message}
			<div
				class="bg-[#86BB71] border rounded-md p-2 m-2 after:content-[' '] after:clear-both after:table"
			>
				<img
					class="h-12 w-12 rounded-full float-left mr-5"
					src="https://images.unsplash.com/photo-1491528323818-fdd1faba62cc?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
					alt=""
				/>
				<p>{message.message}</p>
				<span class="float-right">11:00</span>
			</div>
		{/each}
		<!-- <div
			class="bg-[#94C2ED] border rounded-md p-2 m-2 after:content-[' '] after:clear-both after:table"
		>
			<img
				class="h-12 w-12 rounded-full float-right  mr-0 ml-5"
				src="https://images.unsplash.com/photo-1491528323818-fdd1faba62cc?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
				alt=""
			/>
			<p>Hello. How are you today?</p>
			<span class="float-left">11:00</span>
		</div> -->
	</div>

	<form on:submit|preventDefault={sendMessage} class="p-8">
		<textarea
			class="w-full border-none outline-none py-2 px-5 mb-2 rounded-md resize-none"
			name=""
			id=""
			cols="30"
			rows="10"
			placeholder="Type your message"
			bind:value="{textField}"
		/>
		<button
			class="text-[#94C2ED] bg-[#F2F5F8] hover:text-[#94C2ED] float-right uppercase border-none cursor-pointer font-bold "
			>Send</button
		>
	</form>
</div>

