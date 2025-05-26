<script>
	export let href = '';
	export let text = '';
</script>

<a
	{href}
	class="group relative mt-6 overflow-hidden rounded-xl bg-gradient-to-r from-blue-600 to-indigo-600 px-8 py-3 text-lg font-semibold text-white shadow-lg transition-all duration-300 hover:scale-105"
	target="_blank"
	rel="noopener noreferrer"
>
	{text}
</a>
