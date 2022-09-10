<script lang="ts">
	import { ENDPOINT } from "$lib/realtime";
	import { onMount } from "svelte";

	let textField = "";
	let username =
		Date.now().toString(36) + Math.random().toString(36).substr(2);
	let messages: any[] = [];
	let conn: WebSocket;
	onMount(() => {
		console.log("the component has mounted");
		if (window["WebSocket"]) {
			conn = new WebSocket(
				ENDPOINT +
					"global" +
					"?roomId=global&roomName=global" +
					`&clientId=${username}`
			);
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

		textField = "";
		// Send the message
		conn.send(
			JSON.stringify({
				message: message,
				clientId: username,
				roomId: "global",
			})
		);
	}
</script>

<div>
	<input type="text" placeholder="Username" bind:value={username} />
	<div>
		{#each messages as message}
			{#if message.clientId === username}
				<div>
					<img
						src="https://images.unsplash.com/photo-1491528323818-fdd1faba62cc?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
						alt=""
					/>
					<p>{message.message}</p>
					<span class="float-left">11:00</span>
				</div>
			{:else}
				<div>
					<img
						src="https://images.unsplash.com/photo-1491528323818-fdd1faba62cc?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
						alt=""
					/>
					<p>{message.message}</p>
					<span class="float-right">{new Date().toDateString()}</span>
				</div>
			{/if}
		{/each}
	</div>

	<form on:submit|preventDefault={sendMessage}>
		<textarea
			name=""
			id=""
			cols="30"
			rows="10"
			placeholder="Type your message"
			bind:value={textField}
		/>
		<button>Send</button>
	</form>
</div>
