<script lang="ts">
	import '../styles-reset.scss';
	import IoIosArchive from 'svelte-icons/io/IoIosArchive.svelte';
	import IoIosBarcode from 'svelte-icons/io/IoIosBarcode.svelte';
	import IoIosStats from 'svelte-icons/io/IoIosStats.svelte';
	import IoIosCheckbox from 'svelte-icons/io/IoIosCheckbox.svelte';
	import IoIosLogOut from 'svelte-icons/io/IoIosLogOut.svelte';

	function logout() {
		console.log('logout');
	}
</script>

<div class="layout">
	<nav class="side-menu">
		<div>
			<a href="/profile" class="side-menu__profile">
				<img class="profile-icon" src="src\lib\images\default-user-logo.jpg" alt="" />
				UserName
			</a>
		</div>

		<ul class="side-menu__items">
			<li>
				<a href="/" class="side-menu__item">
					<div class="side-menu__item-icon"><IoIosCheckbox /></div>
					Tasks
				</a>
			</li>
			<li>
				<a href="/habits" class="side-menu__item">
					<div class="side-menu__item-icon"><IoIosBarcode /></div>
					Habits
				</a>
			</li>
			<li>
				<a href="/statistics" class="side-menu__item">
					<div class="side-menu__item-icon"><IoIosStats /></div>
					Statistics
				</a>
			</li>
			<li>
				<a href="/archives" class="side-menu__item">
					<div class="side-menu__item-icon"><IoIosArchive /></div>
					Archives
				</a>
			</li>
		</ul>

		<button class="side-menu__logout" on:click={logout}>
			<div class="side-menu__item-icon"><IoIosLogOut /></div>
			Logout
		</button>
	</nav>
</div>

<slot />

<style lang="scss">
	@import url('https://fonts.googleapis.com/css2?family=Nunito:wght@300;400;600;700&family=Playfair:wght@400;500;700&display=swap');
	:global(body) {
		display: grid;
		grid-template-columns: 300px auto;
		background: #222831;
		color: #f9f9f9;
		font-family: 'Nunito', sans-serif;
		font-size: 14px;
	}

	.side-menu {
		height: 100vh;
		background-color: #2c3440;
		font-size: 18px;
		display: grid;
		grid-template-rows: auto auto 1fr;
		align-items: flex-end;

		&__items {
			padding-top: 15px;
			display: grid;
			row-gap: 30px;
		}

		&__item {
			display: grid;
			grid-template-columns: auto 1fr;
			align-items: center;
			padding-left: 15px;
			column-gap: 10px;
			&-icon {
				width: 24px;
				height: 24px;
			}
		}

		&__profile {
			display: flex;
			align-items: center;
			gap: 10px;
			padding: 15px;
		}

		&__logout {
			display: flex;
			align-items: center;
			gap: 10px;
			padding: 15px;
			cursor: pointer;
		}
	}

	.profile-icon {
		border-radius: 50%;
		width: 40px;
	}
</style>
