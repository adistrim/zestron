<script lang="ts">
    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';
    import { useLayout } from '../../lib/layout';
    import { goto } from '$app/navigation';
    import { page } from '$app/stores';

    let socket: WebSocket | null = null;
    const messages = writable<{ role: string; content: string }[]>([]);
    let input = '';
    let chatId: string | null = null;
    let isGenerating = false;
    let isConnecting = false;
    const layout = useLayout();

    onMount(() => {
        const query = $page.url.search;
        if (query && query.startsWith('?')) {
            const possibleChatId = query.substring(1);
            if (possibleChatId && possibleChatId !== 'undefined' && possibleChatId !== 'null') {
                chatId = possibleChatId;
            }
        }
    });

    function setupWebSocket() {
        if (socket && (socket.readyState === WebSocket.OPEN || socket.readyState === WebSocket.CONNECTING)) {
            return Promise.resolve();
        }

        isConnecting = true;
        return new Promise<void>((resolve, reject) => {
            const API_WS = import.meta.env.VITE_WS_API_BASE_URL + '/ws/generate';
            socket = new WebSocket(API_WS);

            socket.onopen = () => {
                isConnecting = false;
                resolve();
            };

            socket.onerror = () => {
                isConnecting = false;
                isGenerating = false;
                messages.update((m) => [
                    ...m,
                    {
                        role: 'system',
                        content:
                            `❌ Connection error. It's not you, it's us! It's not cheap to keep our WebSocket AI backend running 24/7. If you want to try the platform, connect with us via mail at araj@adistrim.in.`
                    }
                ]);
                reject(new Error('WebSocket connection failed'));
            };

            socket.onmessage = (event) => {
                const data = JSON.parse(event.data);
                isGenerating = false;

                if (data.error) {
                    messages.update((m) => [...m, { role: 'system', content: `❌ ${data.error}` }]);
                    return;
                }

                if (data.message && data.message.content) {
                    messages.update((m) => [...m, { role: 'assistant', content: data.message.content }]);

                    if (data.chatId && data.chatId !== chatId) {
                        chatId = data.chatId;
                        const baseUrl = window.location.pathname;
                        const newUrl = `${baseUrl}?${chatId}`;
                        goto(newUrl, { replaceState: true, noScroll: true });
                    }
                }
            };

            socket.onclose = () => {
                isGenerating = false;
                console.warn('WebSocket closed');
                socket = null;
            };
        });
    }

    async function sendMessage() {
        if (!input.trim() || isGenerating) return;

        try {
            messages.update((m) => [...m, { role: 'user', content: input }]);
            isGenerating = true;
            
            await setupWebSocket();
            
            if (socket && socket.readyState === WebSocket.OPEN) {
                const payload = {
                    chatId: chatId || '',
                    prompt: input
                };

                socket.send(JSON.stringify(payload));
                input = '';
            } else {
                throw new Error('WebSocket not connected');
            }
        } catch (error) {
            isGenerating = false;
            console.error('Failed to send message:', error);
        }
    }

    let chatContainer: HTMLElement;
    $: if ($messages && chatContainer) {
        setTimeout(() => {
            chatContainer.scrollTop = chatContainer.scrollHeight;
        }, 0);
    }
</script>

<svelte:head>
	<title>Create | Zestron</title>
	<meta
		name="description"
		content="Create stunning animations powered by Large Language Models with Zestron"
	/>
</svelte:head>

<section
	class="relative flex flex-col items-center overflow-hidden bg-gradient-to-r from-indigo-50 via-blue-50 to-indigo-50 px-4 sm:px-6"
	style="min-height: calc(100vh - {$layout.headerHeight + $layout.footerHeight}px);"
>
	<div class="z-10 flex w-full max-w-4xl flex-1 flex-col py-6">
		<div class="flex flex-1 flex-col overflow-hidden">
			<div bind:this={chatContainer} class="flex-1 overflow-y-auto p-4 sm:p-6">
				{#each $messages as msg}
					<div
						class={`mb-4 flex ${msg.role === 'user' ? 'justify-end' : msg.role === 'assistant' ? 'justify-start' : 'justify-center'}`}
					>
						<div
							class={`max-w-[85%] rounded-2xl px-4 py-3 ${
								msg.role === 'user'
									? 'bg-gradient-to-r from-blue-600 to-indigo-600 text-white'
									: msg.role === 'assistant'
										? 'bg-gradient-to-r from-blue-500 to-cyan-500 text-white'
										: 'bg-red-100 text-sm text-red-600'
							}`}
						>
							{msg.content}
						</div>
					</div>
				{/each}

				{#if isGenerating}
					<div class="mb-4 flex justify-start">
						<div class="max-w-[85%] rounded-2xl bg-blue-100 px-4 py-3 text-blue-800">
							<div class="flex space-x-2">
								<div class="h-2 w-2 animate-pulse rounded-full bg-blue-400"></div>
								<div
									class="animation-delay-200 h-2 w-2 animate-pulse rounded-full bg-blue-400"
								></div>
								<div
									class="animation-delay-400 h-2 w-2 animate-pulse rounded-full bg-blue-400"
								></div>
							</div>
						</div>
					</div>
				{/if}

				{#if $messages.length === 0 && !isGenerating}
					<div class="flex h-full flex-col items-center justify-center p-6 text-center">
						<div class="mb-4 rounded-full bg-indigo-100 p-3">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								class="text-indigo-600"
							>
								<path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path>
							</svg>
						</div>
						<h2 class="text-xl font-semibold text-gray-800">Get started with Zestron</h2>
						<p class="mt-2 max-w-sm text-sm text-gray-500">
							Describe the animation you want to create, and our AI will help bring your ideas to
							life.
						</p>
					</div>
				{/if}
			</div>

			<div class="rounded-xl bg-white/80 p-3">
				<div class="flex gap-2">
					<input
						bind:value={input}
						type="text"
						placeholder="Describe your animation idea..."
						class="flex-1 rounded-lg border border-gray-200 bg-white px-4 py-3 transition focus:border-indigo-300 focus:outline-none focus:ring-1 focus:ring-indigo-300"
						on:keydown={(e) => e.key === 'Enter' && sendMessage()}
						disabled={isGenerating}
					/>
					<button
						on:click={sendMessage}
						class="rounded-lg bg-gradient-to-r from-blue-600 to-indigo-600 px-4 py-3 font-medium text-white shadow-sm transition hover:from-blue-700 hover:to-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-300 disabled:opacity-50"
						disabled={!input.trim() || isGenerating}
					>
						{#if isGenerating}
							<svg
								class="h-5 w-5 animate-spin"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
							>
								<circle
									class="opacity-25"
									cx="12"
									cy="12"
									r="10"
									stroke="currentColor"
									stroke-width="4"
								></circle>
								<path
									class="opacity-75"
									fill="currentColor"
									d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
								></path>
							</svg>
						{:else}
							Send
						{/if}
					</button>
				</div>
			</div>
		</div>
	</div>
</section>

<style>
	.animation-delay-200 {
		animation-delay: 200ms;
	}
	.animation-delay-400 {
		animation-delay: 400ms;
	}
</style>
