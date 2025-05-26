<script>
	import Button from '../../components/Button.svelte';
	import { useLayout } from '../../lib/layout';

	const layout = useLayout();
</script>

<svelte:head>
	<title>Zestron | AI Animation Generator</title>
	<meta
		name="description"
		content="Create stunning animations powered by Large Language Models with Zestron"
	/>
</svelte:head>

<section
	class="relative flex flex-col items-center justify-center overflow-hidden bg-gradient-to-r from-indigo-50 via-blue-50 to-indigo-50 px-6"
	style="min-height: calc(100vh - {$layout.headerHeight + $layout.footerHeight}px);"
>
	<div class="z-10 flex max-w-2xl flex-col items-center justify-center gap-8 py-12 text-center">
		<div class="mb-4">
			<h1
				class="bg-gradient-to-r from-blue-600 via-indigo-600 to-violet-600 bg-clip-text text-7xl font-extrabold tracking-tight text-transparent"
			>
				Zestron
			</h1>
			<p class="mt-4 text-2xl font-thin tracking-wider text-gray-700">
				Generate Animations using Large Language Models
			</p>
		</div>

		<Button href="/create" text="Get Started" />

		<div class="mt-4">
			<a
				href="https://api-docs.deepseek.com/"
				class="inline-flex items-center gap-2 rounded-full border border-gray-200 bg-white/80 px-5 py-2 text-sm font-medium text-gray-600 shadow-sm backdrop-blur-sm transition-all hover:bg-white hover:shadow"
				target="_blank"
				rel="noopener noreferrer"
			>
				<span>Powered by DeepSeek</span>
				<img
					src="https://brandlogos.net/wp-content/uploads/2025/02/deepseek_logo_icon-logo_brandlogos.net_s5bgc.png"
					alt="DeepSeek Logo"
					class="h-4"
				/>
			</a>
		</div>
	</div>
</section>
