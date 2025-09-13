<script>
  import './assets/css/base.css';
  import Navbar from './lib/components/Navbar.svelte';
  import { onMount } from 'svelte';
  
  // Import your route components
  import HomePage from './lib/routes/HomePage.svelte';
  import SettingsPage from './lib/routes/SettingsPage.svelte';

  // Simple hash-based routing that maintains component state
  let currentRoute = $state('/');

  // Components are rendered once and kept alive
  const components = {
    '/': { component: HomePage, instance: null },
    '/settings': { component: SettingsPage, instance: null }
  };

  // Hash-based routing
  function updateRoute() {
    const hash = window.location.hash.replace('#', '') || '/';
    if (components[hash]) {
      currentRoute = hash;
    }
  }

  onMount(() => {
    // Setup hash routing
    updateRoute();
    window.addEventListener('hashchange', updateRoute);
    
    // Cleanup
    return () => {
      window.removeEventListener('hashchange', updateRoute);
    };
  });
</script>

<div class="app">
  <Navbar {currentRoute} />
  <main>
    {#each Object.entries(components) as [path, { component: Component }]}
      <div class="route-container" class:active={currentRoute === path}>
        <Component />
      </div>
    {/each}
  </main>
</div>

<style>
  .app {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }

  main {
    flex: 1;
    padding: 1.5rem;
    background-color: var(--bg-primary);
    color: var(--text-primary);
    position: relative;
  }

  .route-container {
    display: none;
  }

  .route-container.active {
    display: block;
  }
</style>