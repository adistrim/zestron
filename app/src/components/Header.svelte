<script lang="ts">
	import { page } from '$app/state';
	import { GitBranch } from 'lucide-svelte';

	const url = $derived(page.url);
</script>

<header class="sticky top-0 z-50 bg-gradient-to-r from-indigo-50 via-blue-50 to-indigo-50 border-b border-indigo-100">
	<div class="container mx-auto flex items-center justify-between px-6 py-4">
		<div class="flex items-center">
			<a
				href="/"
				class="bg-gradient-to-r from-blue-600 via-indigo-600 to-violet-600 bg-clip-text text-xl font-bold text-transparent"
			>
				Zestron
			</a>
		</div>

		<nav>
			<ul class="flex items-center space-x-8">
				<li class="relative">
					<a
						href="/"
						class={`text-sm font-medium ${url.pathname === '/' ? 'text-indigo-600' : 'text-gray-700 hover:text-indigo-600'} transition-colors`}
					>
						Home
					</a>
					{#if url.pathname === '/'}
						<div
							class="absolute -bottom-1.5 left-0 right-0 h-0.5 rounded-full bg-gradient-to-r from-blue-600 to-indigo-600"
						></div>
					{/if}
				</li>

				<li class="relative">
					<a
						href="/create"
						class={`text-sm font-medium ${url.pathname === '/create' ? 'text-indigo-600' : 'text-gray-700 hover:text-indigo-600'} transition-colors`}
					>
						Create
					</a>
					{#if url.pathname === '/create'}
						<div
							class="absolute -bottom-1.5 left-0 right-0 h-0.5 rounded-full bg-gradient-to-r from-blue-600 to-indigo-600"
						></div>
					{/if}
				</li>

				<li class="relative">
					<a
						href="/about"
						class={`text-sm font-medium ${url.pathname === '/about' ? 'text-indigo-600' : 'text-gray-700 hover:text-indigo-600'} transition-colors`}
					>
						About
					</a>
					{#if url.pathname === '/about'}
						<div
							class="absolute -bottom-1.5 left-0 right-0 h-0.5 rounded-full bg-gradient-to-r from-blue-600 to-indigo-600"
						></div>
					{/if}
				</li>
			</ul>
		</nav>

		<div>
			<a
				href="https://github.com/adistrim/zestron"
				target="_blank"
				rel="noopener noreferrer"
				class="flex items-center rounded-lg bg-indigo-100 px-4 py-2 text-indigo-600 transition-colors hover:bg-indigo-200 hover:text-indigo-800"
				aria-label="View on GitHub"
			>
				<GitBranch size={18} class="mr-2" />
				GitHub
			</a>
		</div>
	</div>
</header>
