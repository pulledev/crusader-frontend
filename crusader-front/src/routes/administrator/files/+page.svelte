<script lang="ts">
    import UserTable from "$lib/components/UserTable.svelte";
    import { Tooltip } from "bits-ui";
    import QuestionMark from "phosphor-svelte/lib/Question";
    import HelperAccordion from "$lib/components/HelperAccordion.svelte";
    import MouseRightClick from "phosphor-svelte/lib/MouseRightClick";

    // Demo-Daten – in echt aus load() oder API
    const users = [
        { name: "Mia", rang: "OFw.",   element: "Echo",  punkte: 420,  steamId: "123456789", discordId: "1234567890" },
        { name: "Dan", rang: "Hptm.", element: "Alpha", punkte: 1337, steamId: "987654321", discordId: "0987654321" },
        { name: "Paul", rang: "OFw.", element: "Alpha", punkte: 69, steamId: "112233445", discordId: "4206942069" },
        { name: "Pulle", rang: "OFw.", element: "Alpha", punkte: 69, steamId: "123", discordId: "4206942069" },

    ];

    const documentation = [
        {
            title: "Wie öffne ich eine Personalakte?",
            content: "Drücke mit der rechten Maustaste auf ein Tabelleneintrag und wähle Akte öffnen aus"

        },
        {
            title: "Wie bearbeitet man Mitglieder?",
            content: "Drücke mit der rechten Maustaste auf eine Tabelleneintrag und wähle Bearbeiten aus"
        },
        {
            title: "Wie lösche ich Mitglieder?",
            content: "Drücke mit der rechten Maustaste auf eine Tabelleneintrag und wähle Löschen aus"
        }

    ];

    // kleine Stats für Header-Badges
    const total = users.length;
    const uniqueRanks = new Set(users.map(u => u.rang)).size;
    const uniqueElements = new Set(users.map(u => u.element)).size;

    function onCreate() {
        // hier später: goto("/administrator/files/new") o.ä.
        alert("Neuen Datensatz anlegen (Demo)");
    }
    function onExport() {
        // hier später: Export aus deiner Datenquelle
        alert("CSV-Export (Demo)");
    }
</script>

<!-- seichte Fläche & ordentliche Ränder -->
<div class="min-h-[100dvh] bg-gradient-to-b from-slate-50 to-white dark:from-zinc-950 dark:to-zinc-900">
    <main class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 py-8 sm:py-10">

        <!-- Page Header -->
        <header class="mb-6 sm:mb-8">
            <div class="flex items-start justify-between gap-4">
                <div>
                    <h1 class="text-2xl sm:text-3xl font-semibold tracking-tight text-slate-900 dark:text-white">
                        Personalakten
                    </h1>
                    <p class="mt-2 text-sm text-slate-600 dark:text-zinc-400 max-w-prose">
                        Verwalte alle Mitglieder der Kommandokräfte Crusader. Suche filtere und sortiere direkt in der Tabelle.
                    </p>
                </div>

                <!-- Primary Actions -->
                <div class="flex shrink-0 items-center gap-2">
                    <button
                            class="inline-flex items-center gap-2 rounded-xl border border-slate-300 bg-white px-3 py-2 text-sm font-medium text-slate-700
                            hover:bg-slate-50 active:scale-[0.99] dark:border-zinc-700 dark:bg-zinc-900 dark:text-zinc-100 dark:hover:bg-zinc-800"
                            onclick={onExport}
                    >
                        <span aria-hidden="true">↥</span> Export
                    </button>
                    <button
                            class="inline-flex items-center gap-2 rounded-xl bg-slate-900 px-3.5 py-2 text-sm font-semibold text-white  hover:bg-slate-800 active:scale-[0.99] dark:bg-zinc-100 dark:text-zinc-900 dark:hover:bg-zinc-200"
                            onclick={onCreate}
                    >
                        <span aria-hidden="true">＋</span> Neu anlegen
                    </button>
                    <Tooltip.Provider>
                        <Tooltip.Root delayDuration={200}>
                            <Tooltip.Trigger
                                    class="border-border-input bg-background-alt shadow-btn ring-dark ring-offset-background
		hover:bg-muted focus-visible:ring-dark focus-visible:ring-offset-background focus-visible:outline-hidden inline-flex size-10 items-center justify-center rounded-full border focus-visible:ring-2 focus-visible:ring-offset-2"
                                    onclick={() => document.getElementById('documentation-section')?.scrollIntoView({ behavior: 'smooth' })}
                            >
                                <QuestionMark class="size-5" />
                            </Tooltip.Trigger>
                            <Tooltip.Content
                                    sideOffset={8}
                                    class="animate-in fade-in-0 zoom-in-95 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 origin-(--bits-tooltip-content-transform-origin)"
                            >
                                <div
                                        class="rounded-input border-dark-10 bg-background shadow-popover outline-hidden z-0 flex items-center justify-center border p-3 text-sm font-medium"
                                >
                                    zur Dokumentation
                                </div>
                            </Tooltip.Content>
                        </Tooltip.Root>
                    </Tooltip.Provider>
                </div>
            </div>

            <!-- kleine Status-Badges -->
            <div class="mt-4 flex flex-wrap items-center gap-2">
				<span class="inline-flex items-center gap-2 rounded-lg bg-white px-2.5 py-1 text-xs text-slate-600 ring-1 ring-slate-200 dark:bg-zinc-900 dark:text-zinc-300 dark:ring-zinc-800">
					<span class="size-1.5 rounded-full bg-emerald-400/90"></span>
                    Insegsamt {total} Mitglieder
				</span>
            </div>
        </header>

        <!-- Content Card -->
        <section
                class="rounded-2xl border border-slate-200 bg-white/10 p-2 sm:p-3 backdrop-blur dark:border-zinc-800 dark:bg-zinc-950/60"
        >
            <UserTable {users} />
        </section>

        <!-- Documentation Section -->
        <section id="documentation-section" class="mt-8 sm:mt-10">
            <h2 class="mb-4 text-lg font-semibold text-slate-900 dark:text-white">
                Dokumentation & Hilfe
            </h2>
            <div class="rounded-2xl border border-slate-200 bg-white/10 p-4 sm:p-6 backdrop-blur dark:border-zinc-800 dark:bg-zinc-950/60">
                <HelperAccordion
                        items={documentation}
                        type="multiple"
                        collapsible={false}
                />
            </div>
        </section>
    </main>
</div>
