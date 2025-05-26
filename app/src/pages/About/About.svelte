<script>
	import Button from '../../components/Button.svelte';
    import { useLayout } from '../../lib/layout';

    const layout = useLayout();
</script>

<svelte:head>
	<title>About | Zestron</title>
	<meta name="description" content="Learn about Zestron - AI-powered animation generator" />
</svelte:head>

<section
	class="relative flex flex-col items-center justify-center overflow-hidden bg-gradient-to-r from-indigo-50 via-blue-50 to-indigo-50 px-6"
	style="min-height: calc(100vh - {$layout.headerHeight + $layout.footerHeight}px);"
>
	<div class="z-10 flex max-w-2xl flex-col items-center justify-center gap-8 py-12 text-center">
		<div class="mb-8">
			<h1
				class="bg-gradient-to-r from-blue-600 via-indigo-600 to-violet-600 bg-clip-text text-6xl font-extrabold tracking-tight text-transparent"
			>
				About Zestron
			</h1>
		</div>

		<div class="max-w-xl space-y-6 font-thin">
			<p class="text-xl tracking-wider text-gray-700">
				Zestron is an open-source animation tool that turns text prompts into animations using Large
				Language Models. It prioritizes simplicity and is designed for users of all backgrounds.
				Powered by the DeepSeek API, it refines input to produce clear, context-specific visuals for
				creators, educators, and developers alike.
			</p>
		</div>

		<Button href="https://github.com/adistrim/zestron" text="Contribute on GitHub" />
	</div>
</section>
