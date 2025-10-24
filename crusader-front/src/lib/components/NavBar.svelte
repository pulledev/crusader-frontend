<script lang="ts">
    import { NavigationMenu } from "bits-ui";
    import CaretDown from "phosphor-svelte/lib/CaretDown";
    import cn from "clsx";
    import NavBarDropdown from "$lib/components/NavBarDropdown.svelte";

    const components: { title: string; href: string; description: string }[] = [
        {
            title: "Personalakten",
            href: "/administrator/files",
            description:
                "Liste mit allen Mitgliedern"
        },
        {
            title: "Bewerbungsboard",
            href: "/docs/components/alert-dialog",
            description:
                "Ansicht aller neuer Bewerbungen"
        },
        {
            title: "Eventmanager",
            href: "/administrator/events",
            description:
                "Nachbereitung und Anzeige von Events"
        },
        {
            title: "Ausbildungsmanager",
            href: "/docs/components/scroll-area",
            description: "Nachbereitung und Anzeige von Ausbildungen"
        }
    ];

    type ListItemProps = {
        className?: string;
        title: string;
        href: string;
        content: string;
    };
</script>

{#snippet ListItem({ className, title, content, href }: ListItemProps)}
    <li>
        <NavigationMenu.Link
                class={cn(
        "hover:bg-muted hover:text-accent-foreground focus:bg-muted focus:text-accent-foreground outline-hidden block select-none space-y-1 rounded-md p-3 leading-none no-underline transition-colors",
        className
      )}
                {href}
        >
            <div class="text-sm font-medium leading-none">{title}</div>
            <p class="text-muted-foreground line-clamp-2 text-sm leading-snug">
                {content}
            </p>
        </NavigationMenu.Link>
    </li>
{/snippet}

<NavigationMenu.Root class="relative z-10 flex w-full justify-center">
    <NavigationMenu.List
            class="group flex list-none items-center justify-center p-1"
    >
        <NavigationMenu.Item value="getting-started">
            <NavigationMenu.Trigger
                    class="hover:text-accent-foreground focus-visible:bg-muted focus-visible:text-accent-foreground data-[state=open]:shadow-mini dark:hover:bg-muted dark:data-[state=open]:bg-muted focus-visible:outline-hidden group inline-flex h-8 w-max items-center justify-center rounded-[7px] bg-transparent px-4 py-2 text-sm font-medium transition-colors hover:bg-white disabled:pointer-events-none disabled:opacity-50 data-[state=open]:bg-white"
            >
                <span class="hidden sm:inline"> Mitglieds Bereich </span>
                <span class="inline sm:hidden"> Mitglied </span>
                <CaretDown
                        class="relative top-[1px] ml-1 size-3 transition-transform duration-200 group-data-[state=open]:rotate-180"
                        aria-hidden="true"
                />
            </NavigationMenu.Trigger>
            <NavigationMenu.Content
                    class="data-[motion=from-end]:animate-enter-from-right data-[motion=from-start]:animate-enter-from-left data-[motion=to-end]:animate-exit-to-right data-[motion=to-start]:animate-exit-to-left absolute left-0 top-0 w-full sm:w-auto"
            >
                <ul
                        class="m-0 grid list-none gap-x-2.5 p-3 sm:w-[600px] sm:grid-flow-col sm:grid-rows-3 sm:p-[22px]"
                >
                    <li class="row-span-3 mb-2 sm:mb-0">
                        <NavigationMenu.Link
                                href="/dashboard/123456789"
                                class="from-muted/50 to-muted bg-linear-to-b outline-hidden flex h-full w-full select-none flex-col justify-end rounded-md p-6 no-underline focus:shadow-md transition-all duration-500 ease-[cubic-bezier(0.4,0,0.2,1)] brightness-103 hover:shadow-inner hover:brightness-94"
                                style="background-image: url('/backgroun_dashboard.png');
                                background-size: cover;
                                background-position: center;
                                background-repeat: no-repeat;
                                "
                        >
                            <!-- <Icons.logo class="h-6 w-6" /> -->
                            <div class="mb-2 mt-4 text-lg font-medium">Dashboard</div>
                            <p class="text-muted-foreground text-sm leading-tight">
                                Deine Daten auf einem Blick
                            </p>
                        </NavigationMenu.Link>
                    </li>

                    {@render ListItem({
                        href: "/docs",
                        title: "Ausbildungen",
                        content: "Ausbildungskatalog und dein Fortschritt"
                    })}
                    {@render ListItem({
                        href: "/docs/getting-started",
                        title: "Dein Element",
                        content: "Mitlgieder und Rollen in Alpha"
                    })}
                    {@render ListItem({
                        href: "/docs/styling",
                        title: "Events & Ausbildungen",
                        content: "Informationen über alle Events und Ausbildungen"
                    })}
                </ul>
            </NavigationMenu.Content>
        </NavigationMenu.Item>
        <NavigationMenu.Item>
            <NavigationMenu.Trigger
                    class="hover:text-accent-foreground focus-visible:bg-muted focus-visible:text-accent-foreground data-[state=open]:shadow-mini dark:hover:bg-muted dark:data-[state=open]:bg-muted focus-visible:outline-hidden group inline-flex h-8 w-max items-center justify-center rounded-[7px] bg-transparent px-4 py-2 text-sm font-medium transition-colors hover:bg-white disabled:pointer-events-none disabled:opacity-50 data-[state=open]:bg-white"
            >
                <span class="hidden sm:inline"> Stabs Bereich </span>
                <span class="inline sm:hidden"> Stab </span>
                <CaretDown
                        class="relative top-[1px] ml-1 size-3 transition-transform duration-200 group-data-[state=open]:rotate-180"
                        aria-hidden="true"
                />
            </NavigationMenu.Trigger>
            <NavigationMenu.Content
                    class="data-[motion=from-end]:animate-enter-from-right data-[motion=from-start]:animate-enter-from-left data-[motion=to-end]:animate-exit-to-right data-[motion=to-start]:animate-exit-to-left absolute left-0 top-0 w-full sm:w-auto"
            >
                <ul
                        class="grid gap-3 p-3 sm:w-[400px] sm:p-6 md:w-[500px] md:grid-cols-2 lg:w-[600px]"
                >
                    {#each components as component (component.title)}
                        {@render ListItem({
                            href: component.href,
                            title: component.title,
                            content: component.description
                        })}
                    {/each}
                </ul>
            </NavigationMenu.Content>
        </NavigationMenu.Item>
        <NavigationMenu.Item>
            <NavigationMenu.Link
                    class="hover:text-accent-foreground focus:bg-muted focus:text-accent-foreground data-[state=open]:shadow-mini dark:hover:bg-muted dark:data-[state=open]:bg-muted focus:outline-hidden group inline-flex h-8 w-max items-center justify-center rounded-[7px] bg-transparent px-4 py-2 text-sm font-medium transition-colors hover:bg-white disabled:pointer-events-none disabled:opacity-50 data-[state=open]:bg-white"
                    href="/docs"
            >
                <span class="hidden sm:inline"> Administrator Dashboard </span>
                <span class="inline sm:hidden"> Admin Dashboard </span>
            </NavigationMenu.Link>
        </NavigationMenu.Item>
        <NavigationMenu.Item>
            <div class="flex items-center">
                <NavBarDropdown/>
            </div>
        </NavigationMenu.Item>
        <NavigationMenu.Indicator
                class="data-[state=hidden]:animate-fade-out data-[state=visible]:animate-fade-in top-full z-10 flex h-2.5 items-end justify-center overflow-hidden opacity-100 transition-[all,transform_250ms_ease] duration-200 data-[state=hidden]:opacity-0"
        >
            <div
                    class="bg-border relative top-[70%] size-2.5 rotate-[45deg] rounded-tl-[2px]"
            ></div>
        </NavigationMenu.Indicator>
    </NavigationMenu.List>
    <div
            class="perspective-[2000px] absolute left-0 top-full flex w-full justify-center"
    >
        <NavigationMenu.Viewport
                class="text-popover-foreground bg-background data-[state=closed]:animate-scale-out data-[state=open]:animate-scale-in relative mt-2.5 h-[var(--bits-navigation-menu-viewport-height)] w-full origin-[top_center] overflow-hidden rounded-md border shadow-lg transition-[width,_height] duration-200 sm:w-[var(--bits-navigation-menu-viewport-width)] "
        />
    </div>

</NavigationMenu.Root>
