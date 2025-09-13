<script>  
  // Props from parent
  let { currentRoute } = $props();
  
  // Links for HCC dashboard
  let links = [
    { name: 'Home', path: '/' },
    { name: 'Settings', path: '/settings' }
  ];
</script>

<nav class="navbar">
  <h1 class="title">HCC</h1>
  <div class="nav-links">
    {#each links as link}
    <a
        href={`#${link.path}`}
        class:active={currentRoute === link.path}
    >
        {link.name}
    </a>
    {/each}
  </div>
</nav>

<style>
  .navbar {
    display: flex;
    align-items: center;
    gap: 2rem;
    background-color: var(--bg-secondary);
    padding: 1rem;
    border-bottom: 1px solid var(--border-primary);
  }

  .title {
    color: var(--text-primary);
    font-size: 1.5rem;
    font-weight: 600;
    margin: 0;
  }

  .nav-links {
    display: flex;
    gap: 1rem;
  }

  a {
    color: var(--text-primary);
    text-decoration: none;
    padding: 0.5rem 1rem;
    transition: background-color 0.2s ease;
    border-radius: 0.375rem;
    font-weight: 500;
  }

  a:hover {
    background-color: var(--bg-tertiary);
  }

  .active {
    background-color: var(--accent-primary);
    color: white;
  }
</style>