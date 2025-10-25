<script lang="ts">
    import type { PageData } from './$types';
    import HelperAccordion from "$lib/components/HelperAccordion.svelte";
    import { goto } from "$app/navigation";
    import ArrowLeft from "phosphor-svelte/lib/ArrowLeft";
    import { Tooltip } from "bits-ui";
    import QuestionMark from "phosphor-svelte/lib/Question";

    const { data } = $props<{ data: PageData }>();
    const { personnelFile } = data;

    function createEntry() {
        alert("Neuen Eintrag anlegen (Demo)");
    }

    function goBack() {
        goto("/administrator/files");
    }

    // Formatiere Datum
    function formatDate(dateStr: string) {
        return new Date(dateStr).toLocaleDateString("de-DE", {
            day: "2-digit",
            month: "2-digit",
            year: "numeric"
        });
    }

    // Status Badge Farben
    const statusColors = {
        "Aktiv": "bg-emerald-100 text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-400",
        "Inaktiv": "bg-gray-100 text-gray-700 dark:bg-gray-900/30 dark:text-gray-400",
        "Beurlaubt": "bg-amber-100 text-amber-700 dark:bg-amber-900/30 dark:text-amber-400"
    };

    const severityColors = {
        "Leicht": "bg-yellow-100 text-yellow-700 dark:bg-yellow-900/30 dark:text-yellow-400",
        "Mittel": "bg-orange-100 text-orange-700 dark:bg-orange-900/30 dark:text-orange-400",
        "Schwer": "bg-red-100 text-red-700 dark:bg-red-900/30 dark:text-red-400"
    };

    const resultColors = {
        "Erfolg": "bg-emerald-100 text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-400",
        "Teilerfolg": "bg-amber-100 text-amber-700 dark:bg-amber-900/30 dark:text-amber-400",
        "Fehlgeschlagen": "bg-red-100 text-red-700 dark:bg-red-900/30 dark:text-red-400"
    };

    const documentation = [
        {
            title: "Was ist eine Personalakte?",
            content: "Die Personalakte enthält alle relevanten Informationen zu einem Mitglied: Ausbildungsbewertungen, Missionsteilnahmen, Verwarnungen, Beförderungen und Auszeichnungen. Sie dient als zentrale Dokumentation der militärischen Laufbahn."
        },
        {
            title: "Wie lege ich einen neuen Eintrag an?",
            content: "Klicke auf den Button 'Eintrag anlegen' oben rechts. Du kannst dann auswählen, welche Art von Eintrag du hinzufügen möchtest (Ausbildung, Mission, Verwarnung, etc.)."
        },
        {
            title: "Was bedeuten die verschiedenen Status-Badges?",
            content: "Aktiv (grün): Das Mitglied ist derzeit aktiv im Dienst. Inaktiv (grau): Das Mitglied ist derzeit nicht im aktiven Dienst. Beurlaubt (gelb): Das Mitglied ist vorübergehend beurlaubt."
        },
        {
            title: "Wie interpretiere ich Ausbildungsbewertungen?",
            content: "Ausbildungsbewertungen werden mit einer Punktzahl von 0-100 bewertet. Ab 80 Punkten gilt die Ausbildung als sehr gut bestanden, 60-79 als bestanden, unter 60 als nicht bestanden. Die Notizen des Ausbilders geben zusätzlichen Kontext."
        },
        {
            title: "Was bedeuten die Schweregrade bei Verwarnungen?",
            content: "Leicht (gelb): Kleinere Verstöße ohne größere Konsequenzen. Mittel (orange): Bedeutendere Verstöße, können zu Konsequenzen führen. Schwer (rot): Schwerwiegende Verstöße mit ernsthaften Konsequenzen."
        },
        {
            title: "Wie kann ich die Personalakte bearbeiten oder löschen?",
            content: "Verwende den 'Eintrag anlegen' Button, um neue Informationen hinzuzufügen. Das Bearbeiten oder Löschen einzelner Einträge erfolgt über das jeweilige Kontextmenü (Rechtsklick) auf den entsprechenden Eintrag."
        }
    ];
</script>

<div class="min-h-[100dvh] bg-gradient-to-b from-slate-50 to-white dark:from-zinc-950 dark:to-zinc-900">
    <main class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 py-8 sm:py-10">

        <!-- Back Button & Header -->
        <div class="mb-6">
            <button
                    onclick={goBack}
                    class="inline-flex items-center gap-2 text-sm text-slate-600 hover:text-slate-900 dark:text-zinc-400 dark:hover:text-zinc-100 mb-4"
            >
                <ArrowLeft class="size-4" />
                Zurück zur Übersicht
            </button>

            <div class="flex items-start justify-between gap-4 flex-wrap">
                <div>
                    <div class="flex items-center gap-3 mb-2">
                        <h1 class="text-2xl sm:text-3xl font-semibold tracking-tight text-slate-900 dark:text-white">
                            Personalakte: {personnelFile.name}
                        </h1>
                        <span class="inline-flex items-center rounded-lg px-2.5 py-1 text-xs font-medium {statusColors[personnelFile.status]}">
                            {personnelFile.status}
                        </span>
                    </div>
                    <div class="flex flex-wrap gap-x-6 gap-y-2 text-sm text-slate-600 dark:text-zinc-400">
                        <div><span class="font-medium">Rang:</span> {personnelFile.rank}</div>
                        <div><span class="font-medium">Element:</span> {personnelFile.element}</div>
                        <div><span class="font-medium">Beitrittsdatum:</span> {formatDate(personnelFile.joinDate)}</div>
                    </div>
                    <div class="mt-2 flex flex-wrap gap-x-6 gap-y-2 text-xs text-slate-500 dark:text-zinc-500">
                        <div><span class="font-medium">Steam ID:</span> {personnelFile.steamId}</div>
                        <div><span class="font-medium">Discord ID:</span> {personnelFile.discordId}</div>
                    </div>
                </div>

                <!-- Action Buttons -->
                <div class="flex shrink-0 items-center gap-2">
                    <button
                            onclick={createEntry}
                            class="inline-flex items-center gap-2 rounded-xl bg-slate-900 px-3.5 py-2 text-sm font-semibold text-white hover:bg-slate-800 active:scale-[0.99] dark:bg-zinc-100 dark:text-zinc-900 dark:hover:bg-zinc-200"
                    >
                        <span aria-hidden="true">＋</span> Eintrag anlegen
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
                                    Dokumentation
                                </div>
                            </Tooltip.Content>
                        </Tooltip.Root>
                    </Tooltip.Provider>
                </div>
            </div>
        </div>

        <!-- Statistics Cards -->
        <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-4 mb-8">
            <div class="rounded-xl border border-slate-200 bg-white p-4 dark:border-zinc-800 dark:bg-zinc-900">
                <div class="text-2xl font-bold text-slate-900 dark:text-white">{personnelFile.trainingEvaluations.length}</div>
                <div class="text-xs text-slate-600 dark:text-zinc-400">Ausbildungen</div>
            </div>
            <div class="rounded-xl border border-slate-200 bg-white p-4 dark:border-zinc-800 dark:bg-zinc-900">
                <div class="text-2xl font-bold text-slate-900 dark:text-white">{personnelFile.missionParticipations.length}</div>
                <div class="text-xs text-slate-600 dark:text-zinc-400">Missionen</div>
            </div>
            <div class="rounded-xl border border-slate-200 bg-white p-4 dark:border-zinc-800 dark:bg-zinc-900">
                <div class="text-2xl font-bold text-slate-900 dark:text-white">{personnelFile.warnings.length}</div>
                <div class="text-xs text-slate-600 dark:text-zinc-400">Verwarnungen</div>
            </div>
            <div class="rounded-xl border border-slate-200 bg-white p-4 dark:border-zinc-800 dark:bg-zinc-900">
                <div class="text-2xl font-bold text-slate-900 dark:text-white">{personnelFile.promotions.length}</div>
                <div class="text-xs text-slate-600 dark:text-zinc-400">Beförderungen</div>
            </div>
            <div class="rounded-xl border border-slate-200 bg-white p-4 dark:border-zinc-800 dark:bg-zinc-900">
                <div class="text-2xl font-bold text-slate-900 dark:text-white">{personnelFile.awards.length}</div>
                <div class="text-xs text-slate-600 dark:text-zinc-400">Auszeichnungen</div>
            </div>
        </div>

        <!-- Content Sections -->
        <div class="space-y-6">

            <!-- Ausbildungsbewertungen -->
            <section class="rounded-2xl border border-slate-200 bg-white/10 p-4 sm:p-6 backdrop-blur dark:border-zinc-800 dark:bg-zinc-950/60">
                <h2 class="mb-4 text-lg font-semibold text-slate-900 dark:text-white">
                    Ausbildungsbewertungen
                </h2>
                {#if personnelFile.trainingEvaluations.length === 0}
                    <p class="text-sm text-slate-500 dark:text-zinc-400">Keine Ausbildungsbewertungen vorhanden.</p>
                {:else}
                    <div class="space-y-3">
                        {#each personnelFile.trainingEvaluations as training}
                            <div class="rounded-xl border border-slate-200 bg-white p-4 dark:border-zinc-700 dark:bg-zinc-900">
                                <div class="flex items-start justify-between gap-4 mb-2">
                                    <div>
                                        <h3 class="font-medium text-slate-900 dark:text-white">{training.type}</h3>
                                        <p class="text-xs text-slate-500 dark:text-zinc-400">
                                            {formatDate(training.date)} • Ausbilder: {training.trainer}
                                        </p>
                                    </div>
                                    <div class="text-right">
                                        <div class="text-lg font-bold text-slate-900 dark:text-white">{training.score}/100</div>
                                        <div class="text-xs text-slate-500 dark:text-zinc-400">Punkte</div>
                                    </div>
                                </div>
                                <p class="text-sm text-slate-600 dark:text-zinc-300">{training.notes}</p>
                            </div>
                        {/each}
                    </div>
                {/if}
            </section>

            <!-- Missionsteilnahmen -->
            <section class="rounded-2xl border border-slate-200 bg-white/10 p-4 sm:p-6 backdrop-blur dark:border-zinc-800 dark:bg-zinc-950/60">
                <h2 class="mb-4 text-lg font-semibold text-slate-900 dark:text-white">
                    Missionsteilnahmen
                </h2>
                {#if personnelFile.missionParticipations.length === 0}
                    <p class="text-sm text-slate-500 dark:text-zinc-400">Keine Missionsteilnahmen vorhanden.</p>
                {:else}
                    <div class="space-y-3">
                        {#each personnelFile.missionParticipations as mission}
                            <div class="rounded-xl border border-slate-200 bg-white p-4 dark:border-zinc-700 dark:bg-zinc-900">
                                <div class="flex items-start justify-between gap-4">
                                    <div class="flex-1">
                                        <div class="flex items-center gap-2 mb-1">
                                            <h3 class="font-medium text-slate-900 dark:text-white">{mission.mission}</h3>
                                            <span class="inline-flex items-center rounded-md px-2 py-0.5 text-xs font-medium {resultColors[mission.result]}">
                                                {mission.result}
                                            </span>
                                        </div>
                                        <p class="text-xs text-slate-500 dark:text-zinc-400">
                                            {formatDate(mission.date)} • Rolle: {mission.role} • Dauer: {mission.duration}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        {/each}
                    </div>
                {/if}
            </section>

            <!-- Verwarnungen -->
            <section class="rounded-2xl border border-slate-200 bg-white/10 p-4 sm:p-6 backdrop-blur dark:border-zinc-800 dark:bg-zinc-950/60">
                <h2 class="mb-4 text-lg font-semibold text-slate-900 dark:text-white">
                    Verwarnungen
                </h2>
                {#if personnelFile.warnings.length === 0}
                    <p class="text-sm text-slate-500 dark:text-zinc-400">Keine Verwarnungen vorhanden.</p>
                {:else}
                    <div class="space-y-3">
                        {#each personnelFile.warnings as warning}
                            <div class="rounded-xl border border-slate-200 bg-white p-4 dark:border-zinc-700 dark:bg-zinc-900">
                                <div class="flex items-start justify-between gap-4">
                                    <div class="flex-1">
                                        <div class="flex items-center gap-2 mb-1">
                                            <h3 class="font-medium text-slate-900 dark:text-white">{warning.reason}</h3>
                                            <span class="inline-flex items-center rounded-md px-2 py-0.5 text-xs font-medium {severityColors[warning.severity]}">
                                                {warning.severity}
                                            </span>
                                        </div>
                                        <p class="text-xs text-slate-500 dark:text-zinc-400">
                                            {formatDate(warning.date)} • Ausgestellt von: {warning.issuedBy}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        {/each}
                    </div>
                {/if}
            </section>

            <!-- Beförderungen -->
            <section class="rounded-2xl border border-slate-200 bg-white/10 p-4 sm:p-6 backdrop-blur dark:border-zinc-800 dark:bg-zinc-950/60">
                <h2 class="mb-4 text-lg font-semibold text-slate-900 dark:text-white">
                    Beförderungen
                </h2>
                {#if personnelFile.promotions.length === 0}
                    <p class="text-sm text-slate-500 dark:text-zinc-400">Keine Beförderungen vorhanden.</p>
                {:else}
                    <div class="space-y-3">
                        {#each personnelFile.promotions as promotion}
                            <div class="rounded-xl border border-slate-200 bg-white p-4 dark:border-zinc-700 dark:bg-zinc-900">
                                <div class="flex items-start justify-between gap-4">
                                    <div class="flex-1">
                                        <h3 class="font-medium text-slate-900 dark:text-white">
                                            {promotion.fromRank} → {promotion.toRank}
                                        </h3>
                                        <p class="text-xs text-slate-500 dark:text-zinc-400">
                                            {formatDate(promotion.date)} • Befördert von: {promotion.promotedBy}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        {/each}
                    </div>
                {/if}
            </section>

            <!-- Auszeichnungen -->
            <section class="rounded-2xl border border-slate-200 bg-white/10 p-4 sm:p-6 backdrop-blur dark:border-zinc-800 dark:bg-zinc-950/60">
                <h2 class="mb-4 text-lg font-semibold text-slate-900 dark:text-white">
                    Auszeichnungen
                </h2>
                {#if personnelFile.awards.length === 0}
                    <p class="text-sm text-slate-500 dark:text-zinc-400">Keine Auszeichnungen vorhanden.</p>
                {:else}
                    <div class="space-y-3">
                        {#each personnelFile.awards as award}
                            <div class="rounded-xl border border-slate-200 bg-white p-4 dark:border-zinc-700 dark:bg-zinc-900">
                                <div class="flex-1">
                                    <h3 class="font-medium text-slate-900 dark:text-white mb-1">{award.name}</h3>
                                    <p class="text-sm text-slate-600 dark:text-zinc-300 mb-2">{award.description}</p>
                                    <p class="text-xs text-slate-500 dark:text-zinc-400">
                                        {formatDate(award.date)} • Verliehen von: {award.awardedBy}
                                    </p>
                                </div>
                            </div>
                        {/each}
                    </div>
                {/if}
            </section>

        </div>

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
