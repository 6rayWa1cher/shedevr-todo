<script>
	import { page } from '$app/stores';
</script>

<div class="content">
	<h1>{$page.status}: {$page.error?.message}</h1>
</div>
