<script lang="ts">
    import type { Snippet } from "svelte";
    import { ContextMenu, type WithoutChild } from "bits-ui";
    type Props = ContextMenu.Props & {
        trigger: Snippet;
        items: string[];
        contentProps?: WithoutChild<ContextMenu.Content.Props>;
        // other component props if needed
    };
    let {
        open = $bindable(false),
        children,
        trigger,
        items,
        contentProps,
        ...restProps
    }: Props = $props();
</script>

<ContextMenu.Root bind:open {...restProps}>
    <ContextMenu.Trigger>
        {@render trigger()}
    </ContextMenu.Trigger>
    <ContextMenu.Portal>
        <ContextMenu.Content {...contentProps}>
            <ContextMenu.Group>
                <ContextMenu.GroupHeading>Select an Office</ContextMenu.GroupHeading>
                {#each items as item}
                    <ContextMenu.Item textValue={item}>
                        {item}
                    </ContextMenu.Item>
                {/each}
            </ContextMenu.Group>
        </ContextMenu.Content>
    </ContextMenu.Portal>
</ContextMenu.Root>