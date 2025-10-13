<script lang="ts">
    import { DropdownMenu, Avatar } from "bits-ui";

    let mobileOpen = false;

    const navLinks = [
        { href: "/admin", label: "Admin" },
        { href: "/dashboard", label: "User Dashboard" },
        { href: "/users", label: "Users" }
    ];

    function logout() {
        // passe an deinen Go-Handler an
        window.location.href = "/logout";
    }
</script>

<nav class="sticky top-0 z-50 border-b bg-white/80 backdrop-blur supports-[backdrop-filter]:bg-white/60">
    <div class="mx-auto max-w-7xl px-3 sm:px-4">
        <div class="flex h-14 items-center justify-between">

            <!-- Left: Hamburger + Logo -->
            <div class="flex items-center gap-3">
                <button
                        class="inline-flex items-center justify-center rounded-md p-2 hover:bg-gray-100 sm:hidden"
                        aria-label="Navigation öffnen"
                        aria-expanded={mobileOpen}
                        on:click={() => (mobileOpen = !mobileOpen)}
                >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7h16M4 12h16M4 17h16" />
                    </svg>
                </button>

                <a href="/" class="group flex items-center gap-2 font-semibold tracking-tight">
          <span class="grid h-7 w-7 place-items-center rounded-md border bg-white shadow-sm transition group-hover:shadow">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 2l7 4v5c0 5.55-3.84 9.74-7 11-3.16-1.26-7-5.45-7-11V6l7-4z" />
              <path d="M11 7h2v10h-2zM7 11h10v2H7z" class="fill-white" />
            </svg>
          </span>
                    <span class="text-base sm:text-lg">Crusader Portal</span>
                </a>
            </div>

            <!-- Center: Desktop-Navigation -->
            <div class="hidden sm:flex items-center gap-6">
                {#each navLinks as { href, label }}
                    <a href={href} class="text-sm text-gray-700 hover:text-black hover:underline underline-offset-4 transition">
                        {label}
                    </a>
                {/each}
            </div>

            <!-- Right: Account / Dropdown -->
            <div class="flex items-center">
                <DropdownMenu.Root>
                    <DropdownMenu.Trigger class="outline-none">
                        <Avatar.Root class="h-8 w-8 cursor-pointer ring-1 ring-gray-200 rounded-full overflow-hidden">
                            <Avatar.Image src="/avatars/default.png" alt="Account" />
                            <Avatar.Fallback class="text-xs grid place-items-center h-full w-full bg-gray-100">AC</Avatar.Fallback>
                        </Avatar.Root>
                    </DropdownMenu.Trigger>

                    <!-- Portal ist optional; Content kann auch inline gerendert werden -->
                    <DropdownMenu.Portal>
                        <DropdownMenu.Content
                                align="end"
                                class="min-w-44 rounded-md border bg-white p-1.5 shadow-lg focus-visible:outline-none"
                        >
                            <div class="px-2 py-1.5 text-xs text-gray-500">Account</div>

                            <DropdownMenu.Item asChild>
                                <a href="/settings" class="w-full rounded-sm px-2 py-1.5 text-sm hover:bg-gray-50">Einstellungen</a>
                            </DropdownMenu.Item>

                            <DropdownMenu.Separator class="my-1 h-px bg-gray-200" />

                            <DropdownMenu.Item
                                    class="rounded-sm px-2 py-1.5 text-sm hover:bg-gray-50 cursor-pointer"
                                    on:click={logout}
                            >
                                Logout
                            </DropdownMenu.Item>
                        </DropdownMenu.Content>
                    </DropdownMenu.Portal>
                </DropdownMenu.Root>
            </div>
        </div>
    </div>

    <!-- Mobile-Dropdown -->
    {#if mobileOpen}
        <div class="sm:hidden border-t bg-white">
            <div class="mx-auto max-w-7xl px-3 py-3">
                <div class="flex flex-col">
                    {#each navLinks as { href, label }}
                        <a href={href} class="rounded-md px-2 py-2 text-sm hover:bg-gray-50" on:click={() => (mobileOpen = false)}>
                            {label}
                        </a>
                    {/each}
                </div>
            </div>
        </div>
    {/if}
</nav>
