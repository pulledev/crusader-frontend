<script lang="ts">
    import type { PageData } from './$types';

    // Runes: Props reaktiv lesen
    const { data } = $props<{ data: PageData }>();

    // Runes: Abgeleiteter Wert (reaktiv, readonly)
    const id = $derived(data.id);
</script>

<h1>Dashboard von: {id}</h1>
<p>Die ID ist: <strong>{id}</strong></p>

