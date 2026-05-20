<template>
  <div class="relative flex min-h-screen items-center justify-center overflow-hidden bg-dark-50 p-4 dark:bg-dark-950">
    <!-- Brand Background -->
    <div class="absolute inset-0 bg-mesh-gradient bg-[size:auto,auto,56px_56px,56px_56px]"></div>

    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div class="brand-stripe brand-stripe-top"></div>
      <div class="brand-stripe brand-stripe-bottom"></div>
    </div>

    <!-- Content Container -->
    <div class="relative z-10 w-full max-w-md">
      <!-- Logo/Brand -->
      <div class="mb-8 text-center">
        <div
          class="mb-4 inline-flex h-16 w-16 items-center justify-center overflow-hidden rounded-lg border border-dark-100 bg-white shadow-lg shadow-dark-900/10 dark:border-dark-700 dark:bg-dark-900"
        >
          <img :src="siteLogo || '/brand-mark.svg'" alt="Logo" class="h-full w-full object-contain" />
        </div>
        <h1 class="text-gradient mb-2 text-3xl font-bold">
          {{ siteName }}
        </h1>
        <p class="text-sm text-gray-500 dark:text-dark-400">
          {{ siteSubtitle }}
        </p>
      </div>

      <!-- Card Container -->
      <div class="card-glass rounded-lg p-8 shadow-glass">
        <slot />
      </div>

      <!-- Footer Links -->
      <div class="mt-6 text-center text-sm">
        <slot name="footer" />
      </div>

      <!-- Copyright -->
      <div class="mt-8 text-center text-xs text-gray-400 dark:text-dark-500">
        &copy; {{ currentYear }} {{ siteName }}. All rights reserved.
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useAppStore } from '@/stores'
import { sanitizeUrl } from '@/utils/url'

const appStore = useAppStore()

const siteName = computed(() => appStore.siteName || '熔鉴AI')
const siteLogo = computed(() => sanitizeUrl(appStore.siteLogo || '', { allowRelative: true, allowDataUrl: true }))
const siteSubtitle = computed(() => appStore.cachedPublicSettings?.site_subtitle || 'Subscription to API Conversion Platform')
const currentYear = computed(() => new Date().getFullYear())

onMounted(() => {
  appStore.fetchPublicSettings()
})
</script>

<style scoped>
.text-gradient {
  @apply bg-gradient-to-r from-dark-900 via-primary-600 to-accent-500 bg-clip-text text-transparent dark:from-white dark:via-primary-300 dark:to-accent-300;
}

.brand-stripe {
  position: absolute;
  height: 2px;
  background: linear-gradient(90deg, transparent, rgba(7, 159, 211, 0.24), rgba(244, 124, 32, 0.22), transparent);
  transform: rotate(-18deg);
}

.brand-stripe-top {
  left: -8rem;
  top: 9rem;
  width: 36rem;
}

.brand-stripe-bottom {
  right: -6rem;
  bottom: 8rem;
  width: 32rem;
}
</style>
