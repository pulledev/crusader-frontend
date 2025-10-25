<script lang="ts">
    import { Accordion } from "bits-ui";

    export type AccordionItem = {
        title: string;
        content: string;
    };

    // Props
    const {
        items = [] as AccordionItem[],
        type = "single" as "single" | "multiple",
        collapsible = true
    } = $props<{
        items?: AccordionItem[];
        type?: "single" | "multiple";
        collapsible?: boolean;
    }>();


    let value = $state<string | string[] | null>(type === "single" ? null : []);

    $effect(() => {
        value = type === "single" ? null : [];
    });

    const idFor = (i: number) => `item-${i}`;
</script>

<div class="w-full">
    <Accordion.Root
            type={type}
            bind:value={value}
            {collapsible}
            class="space-y-3"
    >
        {#each items as item, i (idFor(i))}
            <Accordion.Item
                    value={idFor(i)}
                    class="rounded-xl border border-slate-200 dark:border-zinc-700 overflow-hidden bg-white dark:bg-zinc-900"
            >
                <Accordion.Trigger
                        class="flex w-full items-center justify-between gap-2 px-4 py-3.5 text-left text-sm font-medium text-slate-800 hover:bg-slate-50 focus:outline-none transition-colors dark:text-zinc-100 dark:hover:bg-zinc-800"
                >
                    <span>{item.title}</span>
                    <span
                            class="text-lg transition-transform duration-200 data-[state=open]:rotate-180 text-slate-500 dark:text-zinc-400"
                            aria-hidden="true"
                    >⌄</span>
                </Accordion.Trigger>

                <Accordion.Content
                        class="px-4 py-3 text-sm text-slate-600 dark:text-zinc-300 border-t border-slate-100 dark:border-zinc-800 bg-slate-50/50 dark:bg-zinc-900/50"
                >
                    {item.content}
                </Accordion.Content>
            </Accordion.Item>
        {/each}
    </Accordion.Root>
</div>
