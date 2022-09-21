<script lang="ts">
	import { ENDPOINT } from "$lib/realtime";
	import { onMount } from "svelte";
	import "../../app.css";
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

		const message = textField.trim();

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

<div class="flex flex-row h-full w-full overflow-x-hidden">
	<div class="flex flex-col py-8 pl-6 pr-2 w-64 bg-white flex-shrink-0">
		<div class="flex flex-row items-center justify-center h-12 w-full">
			<div
				class="flex items-center justify-center rounded-2xl text-indigo-700 bg-indigo-100 h-10 w-10"
			>
				<svg
					class="w-6 h-6"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"
					/>
				</svg>
			</div>
			<div class="ml-2 font-bold text-2xl">QuickChat</div>
		</div>
		<div
			class="flex flex-col items-center bg-indigo-100 border border-gray-200 mt-4 w-full py-6 px-4 rounded-lg"
		>
			<div class="h-20 w-20 rounded-full border overflow-hidden">
				<img
					src="https://robohash.org/ab6ef9bb0974d105efd6d0cab1eec20a?set=set2&bgset=bg2&size=200x200"
					alt="Avatar"
					class="h-full w-full"
				/>
			</div>
			<div class="text-sm font-semibold mt-2">Aminos Co.</div>
			<div class="text-xs text-gray-500">Lead UI/UX Designer</div>
			<div class="flex flex-row items-center mt-3">
				<div
					class="flex flex-col justify-center h-4 w-8 bg-indigo-500 rounded-full"
				>
					<div class="h-3 w-3 bg-white rounded-full self-end mr-1" />
				</div>
				<div class="leading-none ml-1 text-xs">Active</div>
			</div>
		</div>
		<div class="flex flex-col mt-8">
			<div class="flex flex-row items-center justify-between text-xs">
				<span class="font-bold">Active Conversations</span>
				<span
					class="flex items-center justify-center bg-gray-300 h-4 w-4 rounded-full"
					>4</span
				>
			</div>
			<div class="flex flex-col space-y-1 mt-4 -mx-2 h-48 overflow-auto">
				<a href="/room">
					<div
						class="flex items-center hover:bg-gray-100 rounded-xl gap-4 p-4"
					>
						<div
							class="flex items-center justify-center h-8 w-8 bg-indigo-200 rounded-full"
						>
							H
						</div>
						<div class="flex flex-col">
							<strong class="text-slate-900 text-sm font-medium"
								>Andrew Alfred</strong
							>
							<span class="text-slate-500 text-sm font-medium"
								>Technical advisor</span
							>
						</div>
					</div>
				</a>

				<div
					class="flex items-center hover:bg-gray-100 rounded-xl gap-4 p-4"
				>
					<div
						class="flex items-center justify-center h-8 w-8 bg-indigo-200 rounded-full"
					>
						H
					</div>
					<div class="flex flex-col">
						<strong class="text-slate-900 text-sm font-medium"
							>Andrew Alfred</strong
						>
						<span class="text-slate-500 text-sm font-medium"
							>Technical advisor</span
						>
					</div>
				</div>
			</div>
		</div>
	</div>
	<div class="flex flex-col flex-auto h-full p-6">
		<div
			class="flex flex-col flex-auto flex-shrink-0 rounded-2xl bg-gray-100 h-full p-4"
		>
			<div class="flex flex-col h-full overflow-auto mb-4">
				<!-- History chat -->
				<div class="flex flex-col h-full">
					<div class="grid grid-cols-12 gap-y-2">
						{#each messages as message}
							{#if message.clientId === username}
								<div
									class="col-start-6 col-end-13 p-3 rounded-lg"
								>
									<div
										class="flex items-center justify-start flex-row-reverse"
									>
										<div
											class="flex items-center justify-center h-10 w-10 rounded-full bg-indigo-500 flex-shrink-0"
										>
											{username.substring(0, 1)}
										</div>
										<div
											class="relative mr-3 text-sm bg-indigo-100 py-2 px-4 shadow rounded-xl"
										>
											<div>{message.message}</div>
										</div>
									</div>
								</div>
							{:else}
								<div
									class="col-start-1 col-end-8 p-3 rounded-lg"
								>
									<div class="flex flex-row items-center">
										<div
											class="flex items-center justify-center h-10 w-10 rounded-full bg-indigo-500 flex-shrink-0"
										>
											{username.substring(0, 1)}
										</div>
										<div
											class="relative ml-3 text-sm bg-white py-2 px-4 shadow rounded-xl"
										>
											<div>{message.message}</div>
										</div>
									</div>
								</div>
							{/if}
						{/each}
					</div>
				</div>
			</div>
			<!-- Chat -->
			<div
				class="flex flex-row items-center h-16 rounded-xl bg-white w-full px-4"
			>
				<div>
					<button
						class="flex items-center justify-center text-gray-400 hover:text-gray-600"
					>
						<svg
							class="w-5 h-5"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
							xmlns="http://www.w3.org/2000/svg"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13"
							/>
						</svg>
					</button>
				</div>
				<div class="flex-grow ml-4">
					<div class="relative w-full">
						<input
							type="text"
							bind:value={textField}
							class="flex w-full border rounded-xl focus:outline-none focus:border-indigo-300 pl-4 h-10"
						/>
						<button
							class="absolute flex items-center justify-center h-full w-12 right-0 top-0 text-gray-400 hover:text-gray-600"
						>
							<svg
								class="w-6 h-6"
								fill="none"
								stroke="currentColor"
								viewBox="0 0 24 24"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
								/>
							</svg>
						</button>
					</div>
				</div>

				<div class="ml-4">
					<form on:submit|preventDefault={sendMessage}>
						<button
							class="flex items-center justify-center bg-indigo-500 hover:bg-indigo-600 rounded-xl text-white px-4 py-1 flex-shrink-0"
						>
							<span>Send</span>
							<span class="ml-2">
								<svg
									class="w-4 h-4 transform rotate-45 -mt-px"
									fill="none"
									stroke="currentColor"
									viewBox="0 0 24 24"
									xmlns="http://www.w3.org/2000/svg"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
									/>
								</svg>
							</span>
						</button>
					</form>
				</div>
			</div>
		</div>
	</div>
</div>
