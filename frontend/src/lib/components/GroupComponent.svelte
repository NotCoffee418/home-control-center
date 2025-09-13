<script>
  export let title = '';
</script>

<div class="group">
  {#if title}
    <h2 class="group-title">{title}</h2>
  {/if}
  <slot></slot>
</div>

<style>
  .group {
    background-color: #1f2937;
    border-radius: 0.5rem;
    border: 1px solid #374151;
    box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
    padding: 1.5rem;
  }

  .group-title {
    font-size: 1.125rem;
    font-weight: 600;
    color: #f9fafb;
    margin: 0 0 1rem 0;
    padding-bottom: 0.5rem;
    border-bottom: 1px solid #4b5563;
  }
</style>