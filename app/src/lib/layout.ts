import { writable } from 'svelte/store';
import { onMount, onDestroy } from 'svelte';

interface Layout {
    headerHeight: number;
    footerHeight: number;
}

export function useLayout() {
    const layout = writable<Layout>({
        headerHeight: 0,
        footerHeight: 0,
    });

    const updateHeights = () => {
        const header = document.querySelector('header');
        const footer = document.querySelector('footer');

        layout.update(() => ({
            headerHeight: header ? header.offsetHeight : 0,
            footerHeight: footer ? footer.offsetHeight : 0,
        }));
    };

    onMount(() => {
        updateHeights();
        window.addEventListener('resize', updateHeights);

        onDestroy(() => {
            window.removeEventListener('resize', updateHeights);
        });
    });

    return layout;
}
