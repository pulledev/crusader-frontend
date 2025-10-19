<script lang="ts">
    import { Avatar, DropdownMenu } from "bits-ui";
    import Cardholder from "phosphor-svelte/lib/Cardholder";
    import CaretRight from "phosphor-svelte/lib/CaretRight";
    import DotsThree from "phosphor-svelte/lib/DotsThree";
    import GearSix from "phosphor-svelte/lib/GearSix";
    import SignOut from "phosphor-svelte/lib/SignOut";
    import UserCircle from "phosphor-svelte/lib/UserCircle";
    import UserCirclePlus from "phosphor-svelte/lib/UserCirclePlus";
    import Bell from "phosphor-svelte/lib/Bell";
    import Check from "phosphor-svelte/lib/Check";
    import DotOutline from "phosphor-svelte/lib/DotOutline";

    let notifications = $state<boolean>(false);
    let invited = $state("");
</script>

<DropdownMenu.Root>
    <DropdownMenu.Trigger
            class="border-input text-foreground  hover:bg-muted inline-flex h-10 w-10 select-none items-center justify-center rounded-full text-sm font-medium active:scale-[0.98]"
    >
        <DotsThree class="text-foreground h-6 w-6" />
    </DropdownMenu.Trigger>
    <DropdownMenu.Portal>
        <DropdownMenu.Content
                class="border-muted bg-background shadow-popover outline-hidden focus-visible:outline-hidden w-[229px] rounded-xl border px-1 py-1.5"
                sideOffset={8}
        >
            <DropdownMenu.Item
                    class="rounded-button data-highlighted:bg-muted ring-0! ring-transparent! flex h-10 select-none items-center py-3 pl-3 pr-1.5 text-sm font-medium focus-visible:outline-none"
            >
                <div class="flex items-center">
                    <UserCircle class="text-foreground-alt mr-2 size-5" />
                    Profil
                </div>
            </DropdownMenu.Item>
            <DropdownMenu.Item
                    class="rounded-button data-highlighted:bg-muted ring-0! ring-transparent! flex h-10 select-none items-center py-3 pl-3 pr-1.5 text-sm font-medium focus-visible:outline-none"
            >
                <div class="flex items-center">
                    <GearSix class="text-foreground-alt mr-2 size-5" />
                    Einstellungen
                </div>
            </DropdownMenu.Item>
            <DropdownMenu.Item
                    class="rounded-button data-highlighted:bg-muted ring-0! ring-transparent! flex h-10 select-none items-center py-3 pl-3 pr-1.5 text-sm font-medium focus-visible:outline-none"
            >
                <div class="flex items-center">
                    <SignOut class="text-foreground-alt mr-2 size-5" />
                    Abmelden
                </div>
            </DropdownMenu.Item>
        </DropdownMenu.Content>
    </DropdownMenu.Portal>
</DropdownMenu.Root>