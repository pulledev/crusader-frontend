<script lang="ts">
    import { ContextMenu } from "bits-ui";
    import { goto } from "$app/navigation";
    import Mouse from "phosphor-svelte/lib/Mouse";
    import Edit from "phosphor-svelte/lib/PencilSimple";
    import Delete from "phosphor-svelte/lib/Trash";

    type User = {
        name: string;
        rang: string;
        element: string;
        punkte: number;
        steamId: string;
        discordId: string;
        mitgliedstyp: "Mitglied" | "Inaktiv" | "Freund";
    };

    const { users = [] as User[] } = $props<{ users?: User[] }>();

    let search = $state("");
    let filterRang = $state<string>("alle");
    let filterElement = $state<string>("alle");
    let showOnlyActive = $state(false);
    let pageSize = $state(10);
    let page = $state(1);

    type SortKey = keyof User;
    let sortKey = $state<SortKey>("name");
    let sortDir = $state<"asc" | "desc">("asc");

    function toggleSort(key: SortKey) {
        if (sortKey === key) {
            sortDir = sortDir === "asc" ? "desc" : "asc";
        } else {
            sortKey = key;
            sortDir = "asc";
        }
        page = 1;
    }

    const normalized = (v: unknown) =>
        String(v ?? "")
            .toLowerCase()
            .normalize("NFKD")
            .replace(/\p{Diacritic}/gu, "");

    function cmp(a: unknown, b: unknown) {
        const na = normalized(a);
        const nb = normalized(b);
        if (!isNaN(Number(a)) && !isNaN(Number(b))) {
            return Number(a) - Number(b);
        }
        return na < nb ? -1 : na > nb ? 1 : 0;
    }

    const ranks = $derived(
        Array.from(new Set(users.map((u) => u.rang))).sort((a, b) => cmp(a, b))
    );
    const elements = $derived(
        Array.from(new Set(users.map((u) => u.element))).sort((a, b) => cmp(a, b))
    );

    // ✅ KORREKT: Callback-Form mit $derived.by
    const filtered = $derived.by(() => {
        const q = normalized(search);
        return users.filter((u) => {
            const matchSearch =
                !q ||
                [u.name, u.rang, u.element, u.punkte, u.steamId, u.discordId, u.mitgliedstyp]
                    .map((v) => normalized(v))
                    .some((s) => s.includes(q));
            const matchRang = filterRang === "alle" || u.rang === filterRang;
            const matchElement = filterElement === "alle" || u.element === filterElement;
            const matchActive = !showOnlyActive || (u.mitgliedstyp === "Mitglied");
            return matchSearch && matchRang && matchElement && matchActive;
        });
    });

    // darf Ausdruck bleiben
    const sorted = $derived(
        [...filtered].sort((a, b) => {
            const c = cmp(a[sortKey], b[sortKey]);
            return sortDir === "asc" ? c : -c;
        })
    );

    const total = $derived(sorted.length);
    const totalPages = $derived(Math.max(1, Math.ceil(total / pageSize)));

    $effect(() => {
        if (page > totalPages) page = totalPages;
        if (page < 1) page = 1;
    });

    // ✅ KORREKT: auch hier $derived.by
    const pageSlice = $derived.by(() => {
        const start = (page - 1) * pageSize;
        return sorted.slice(start, start + pageSize);
    });

    function openAkte(u: User) {
        goto(`/administrator/files/${u.steamId}`);
    }
    function editUser(u: User) {
        alert(`Bearbeiten: ${u.name}`);
    }
    function deleteUser(u: User) {
        const ok = confirm(`Soll ${u.name} wirklich gelöscht werden?`);
        if (ok) {
            alert(`(Demo) ${u.name} gelöscht`);
        }
    }

    $effect(() => {
        if (users.length === 0) {
            console.warn("UserTable ohne Daten gerendert – Demo-Modus aktiv.");
        }
    });
</script>

<!-- Toolbar -->
<div class="mb-4 flex flex-col gap-3 md:flex-row md:items-end md:justify-between">
    <div class="flex flex-col gap-2 sm:flex-row sm:items-center">
        <div class="relative">
            <input
                    class="w-72 rounded-xl border border-gray-300 bg-white px-3 py-2 text-sm outline-none ring-0 placeholder:text-gray-400 focus:border-gray-400 focus:ring-2 focus:ring-gray-200 dark:border-zinc-700 dark:bg-zinc-900 dark:focus:ring-zinc-700"
                    placeholder="Suche (Name, IDs, Rang, Element)…"
                    bind:value={search}
            />
            {#if search}
                <button
                        class="absolute right-2 top-1/2 -translate-y-1/2 rounded-md px-2 text-xs text-gray-500 hover:text-gray-700 dark:text-zinc-400 dark:hover:text-zinc-200"
                        onclick={() => (search = "")}
                        aria-label="Suche leeren"
                >
                    ✕
                </button>
            {/if}
        </div>

        <!-- Nur Aktive Checkbox -->
        <label class="flex items-center gap-2 text-sm text-gray-700 dark:text-zinc-300 cursor-pointer">
            <input
                    type="checkbox"
                    bind:checked={showOnlyActive}
                    class="size-4 rounded border-gray-300 text-slate-900 focus:ring-2 focus:ring-slate-400 dark:border-zinc-600 dark:bg-zinc-800 dark:text-zinc-100 dark:focus:ring-zinc-500"
            />
            <span>Nur Aktive</span>
        </label>
    </div>
    <div class="mt-4 flex flex-col items-center justify-between gap-3 sm:flex-row">
        <div class="text-sm text-gray-500 dark:text-zinc-400 absolut"> {total} Einträge </div>
    </div>
</div>

<div class="overflow-x-auto">
    <table class="min-w-full text-left text-sm">
        <thead class="bg-gray-50 text-xs uppercase tracking-wide text-gray-500 dark:bg-zinc-900 dark:text-zinc-400">
        <tr>
            {#each [
                { key: "name", label: "Name" },
                { key: "rang", label: "Rang" },
                { key: "element", label: "Element" },
                { key: "mitgliedstyp", label: "Mitgliedstyp" },
                { key: "punkte", label: "Punkte" },
                { key: "steamId", label: "SteamId" },
                { key: "discordId", label: "DiscordId" }
            ] as col}
                <th class="px-4 py-3">
                    <button
                            class="group inline-flex items-center gap-1 font-medium text-gray-600 hover:text-gray-900 dark:text-zinc-300 dark:hover:text-white"
                            onclick={() => toggleSort(col.key as SortKey)}
                    >
                        {col.label}
                        <span
                                class="select-none text-[10px] opacity-0 transition-opacity group-hover:opacity-100"
                                aria-hidden="true"
                        >
								{sortKey === col.key ? (sortDir === "asc" ? "▲" : "▼") : "⇅"}
							</span>
                    </button>
                </th>
            {/each}
        </tr>
        </thead>

        <tbody class="divide-y divide-gray-100 dark:divide-zinc-800">
        {#if pageSlice.length === 0}
            <tr>
                <td colspan="7" class="px-4 py-10 text-center text-gray-500 dark:text-zinc-400">
                    Keine Treffer.
                </td>
            </tr>
        {:else}
            {#each pageSlice as u (u.steamId)}
                <ContextMenu.Root>
                    <ContextMenu.Trigger>
                        {#snippet child({ props })}
                            <tr {...props} class="hover:bg-gray-50 dark:hover:bg-zinc-900/50 cursor-default">
                                <td class="px-4 py-3 font-medium text-gray-900 dark:text-white">{u.name}</td>
                                <td class="px-4 py-3">{u.rang}</td>
                                <td class="px-4 py-3">{u.element}</td>
                                <td class="px-4 py-3">
                                    <span class="inline-flex items-center rounded-md px-2 py-0.5 text-xs font-medium
                                        {u.mitgliedstyp === 'Mitglied' ? 'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-400' :
                                         u.mitgliedstyp === 'Freund' ? 'bg-blue-100 text-blue-700 dark:bg-blue-900/30 dark:text-blue-400' :
                                         'bg-gray-100 text-gray-700 dark:bg-gray-900/30 dark:text-gray-400'}">
                                        {u.mitgliedstyp}
                                    </span>
                                </td>
                                <td class="px-4 py-3 tabular-nums">{u.punkte}</td>
                                <td class="px-4 py-3 font-mono text-xs">{u.steamId}</td>
                                <td class="px-4 py-3 font-mono text-xs">{u.discordId}</td>
                                <td>

                                </td>
                            </tr>
                        {/snippet}
                    </ContextMenu.Trigger>

                    <ContextMenu.Portal>
                        <ContextMenu.Content
                                class="z-50 w-48 rounded-xl border border-gray-200 bg-white p-1 shadow-lg outline-none dark:border-zinc-700 dark:bg-zinc-900"
                        >
                            <ContextMenu.Item
                                    class="flex h-9 select-none items-center rounded-lg px-3 text-sm hover:bg-gray-100 focus:bg-gray-100 dark:hover:bg-zinc-800 dark:focus:bg-zinc-800"
                                    onclick={() => openAkte(u)}
                            >
                                <div class="flex h-9 select-none items-center gap-2 rounded-lg px-3 text-sm hover:bg-gray-100 focus:bg-gray-100 dark:hover:bg-zinc-800 dark:focus:bg-zinc-800">
                                    <Mouse class="size-4.5"/>Akte öffnen
                                </div>
                            </ContextMenu.Item>
                            <ContextMenu.Item
                                    class="flex h-9 select-none items-center rounded-lg px-3 text-sm hover:bg-gray-100 focus:bg-gray-100 dark:hover:bg-zinc-800 dark:focus:bg-zinc-800"
                                    onclick={() => editUser(u)}
                            >
                                <div class="flex h-9 select-none items-center gap-2 rounded-lg px-3 text-sm hover:bg-gray-100 focus:bg-gray-100 dark:hover:bg-zinc-800 dark:focus:bg-zinc-800">
                                    <Edit class="size-4.5"/>Bearbeiten
                                </div>
                            </ContextMenu.Item>
                            <ContextMenu.Item
                                    class="flex h-9 select-none items-center rounded-lg px-3 text-sm text-red-600 hover:bg-red-50 focus:bg-red-50 dark:hover:bg-red-950/30 dark:focus:bg-red-950/30"
                                    onclick={() => deleteUser(u)}
                            >
                                <div class="flex h-9 select-none items-center gap-2 rounded-lg px-3 text-sm  focus:bg-gray-100 dark:hover:bg-zinc-800 dark:focus:bg-zinc-800">

                                <Delete class="size-4.5"/>Löschen

                                </div>

                            </ContextMenu.Item>
                        </ContextMenu.Content>
                    </ContextMenu.Portal>
                </ContextMenu.Root>
            {/each}
        {/if}
        </tbody>
    </table>
</div>



<style>
    th button { width: 100%; text-align: left; }
</style>
