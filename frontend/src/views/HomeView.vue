<template>
  <!-- Custom Home Content: Full Page Mode -->
  <div v-if="homeContent" class="min-h-screen">
    <!-- iframe mode -->
    <iframe
      v-if="isHomeContentUrl"
      :src="homeContent.trim()"
      class="h-screen w-full border-0"
      allowfullscreen
    ></iframe>
    <div v-else v-html="sanitizedHomeContent"></div>
  </div>

  <!-- Default Home Page -->
  <div
    v-else
    class="brand-home relative flex min-h-screen flex-col overflow-hidden bg-dark-50 text-dark-900 dark:bg-dark-950 dark:text-white"
  >
    <!-- Brand Background -->
    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div class="absolute inset-0 bg-mesh-gradient bg-[size:auto,auto,56px_56px,56px_56px]"></div>
      <div class="brand-diagonal brand-diagonal-top"></div>
      <div class="brand-diagonal brand-diagonal-bottom"></div>
    </div>

    <!-- Header -->
    <header
      class="sticky top-0 z-30 border-b border-white/5 bg-dark-950/82 px-4 py-4 backdrop-blur-xl sm:px-6"
    >
      <nav class="mx-auto flex min-w-0 max-w-6xl items-center justify-between gap-3">
        <router-link to="/" class="flex min-w-0 shrink-0 items-center gap-3 transition-opacity hover:opacity-85">
          <div
            class="h-10 w-10 overflow-hidden rounded-lg border border-dark-100 bg-white shadow-sm dark:border-dark-700 dark:bg-dark-900"
          >
            <img :src="siteLogo || '/brand-mark.svg'" alt="熔鉴AI" class="h-full w-full object-contain" />
          </div>
          <div class="min-w-0">
            <div class="text-lg font-bold leading-tight text-white">
              熔鉴<span class="text-primary-400">AI</span>
            </div>
            <p class="hidden text-[10px] leading-4 text-dark-500 sm:block">
              {{ t('home.brandTagline') }}
            </p>
          </div>
        </router-link>

        <!-- Nav Actions -->
        <div class="flex min-w-0 shrink-0 items-center gap-1.5 sm:gap-3">
          <router-link
            :to="isAuthenticated ? '/usage' : '/login'"
            class="hidden whitespace-nowrap text-xs font-medium text-dark-300 transition-colors hover:text-primary-400 sm:inline-flex"
          >
            {{ t('home.checkUsage') }}
          </router-link>
          <router-link
            :to="isAuthenticated ? '/purchase' : '/login'"
            class="hidden whitespace-nowrap rounded-lg bg-primary-500/10 px-3 py-1.5 text-xs font-semibold text-primary-400 transition-colors hover:bg-primary-500/20 sm:inline-flex"
          >
            {{ t('home.buyPlans') }}
          </router-link>

          <!-- Language Switcher -->
          <LocaleSwitcher />

          <!-- Doc Link -->
          <a
            v-if="docUrl"
            :href="docUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 hover:text-gray-700 dark:text-dark-400 dark:hover:bg-dark-800 dark:hover:text-white"
            :title="t('home.viewDocs')"
          >
            <Icon name="book" size="md" />
          </a>

          <!-- Theme Toggle -->
          <button
            @click="toggleTheme"
            class="rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 hover:text-gray-700 dark:text-dark-400 dark:hover:bg-dark-800 dark:hover:text-white"
            :title="isDark ? t('home.switchToLight') : t('home.switchToDark')"
          >
            <Icon v-if="isDark" name="sun" size="md" />
            <Icon v-else name="moon" size="md" />
          </button>

          <!-- Login / Dashboard Button -->
          <router-link
            v-if="isAuthenticated"
            :to="dashboardPath"
            class="inline-flex items-center gap-1.5 rounded-lg bg-dark-900 py-1 pl-1 pr-2.5 transition-colors hover:bg-dark-800 dark:bg-white dark:text-dark-950 dark:hover:bg-dark-100"
          >
            <span
              class="flex h-5 w-5 items-center justify-center rounded-md bg-gradient-accent text-[10px] font-semibold text-white"
            >
              {{ userInitial }}
            </span>
            <span class="text-xs font-medium text-white dark:text-dark-950">{{ t('home.dashboard') }}</span>
            <svg
              class="h-3 w-3 text-gray-400"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M4.5 19.5l15-15m0 0H8.25m11.25 0v11.25"
              />
            </svg>
          </router-link>
          <router-link
            v-else
            to="/login"
            class="inline-flex items-center rounded-lg bg-dark-900 px-3 py-1.5 text-xs font-medium text-white transition-colors hover:bg-dark-800 dark:bg-white dark:text-dark-950 dark:hover:bg-dark-100"
          >
            {{ t('home.login') }}
          </router-link>
        </div>
      </nav>
    </header>

    <div class="sticky top-[73px] z-20 border-b border-white/5 bg-dark-950/90 px-4 py-3 backdrop-blur-lg sm:hidden">
      <div class="flex gap-2 overflow-x-auto">
        <router-link
          :to="isAuthenticated ? '/purchase' : '/login'"
          class="shrink-0 rounded-full border border-white/10 bg-white/5 px-3 py-1.5 text-xs font-medium text-dark-200 transition-colors hover:bg-white/10"
        >
          {{ t('home.buyPlans') }}
        </router-link>
        <router-link
          :to="isAuthenticated ? '/usage' : '/login'"
          class="shrink-0 rounded-full border border-white/10 bg-white/5 px-3 py-1.5 text-xs font-medium text-dark-200 transition-colors hover:bg-white/10"
        >
          {{ t('home.checkUsage') }}
        </router-link>
        <a
          v-if="docUrl"
          :href="docUrl"
          target="_blank"
          rel="noopener noreferrer"
          class="shrink-0 rounded-full border border-white/10 bg-white/5 px-3 py-1.5 text-xs font-medium text-dark-200 transition-colors hover:bg-white/10"
        >
          {{ t('home.docs') }}
        </a>
      </div>
    </div>

    <!-- Main Content -->
    <main class="relative z-10 flex-1 px-4 py-12 sm:px-6">
      <div class="mx-auto min-w-0 max-w-6xl">
        <!-- Hero Section -->
        <div class="mb-12 flex flex-col items-center gap-10 text-center">
          <div class="mx-auto max-w-4xl">
            <div class="mb-6 inline-flex items-center gap-2 rounded-full border border-primary-500/20 bg-primary-500/10 px-4 py-2 text-sm text-primary-300">
              <Icon name="shield" size="sm" />
              {{ t('home.heroBadge') }}
            </div>
            <h1
              class="hero-gradient-text mb-4 text-5xl font-extrabold leading-tight md:text-6xl lg:text-7xl"
            >
              熔鉴AI
            </h1>
            <p class="mb-3 text-2xl font-semibold text-white md:text-3xl">
              {{ t('home.heroHeadline') }}
            </p>
            <p class="mx-auto mb-7 max-w-2xl break-words text-base leading-7 text-dark-400 md:text-lg">
              {{ heroSubtitle }}
            </p>

            <div class="flex flex-col items-center justify-center gap-3 sm:flex-row">
              <router-link
                :to="isAuthenticated ? '/purchase' : '/login'"
                class="inline-flex items-center justify-center rounded-xl bg-gradient-accent px-6 py-3 text-base font-semibold text-white shadow-lg shadow-accent-500/20 transition-transform hover:-translate-y-0.5"
              >
                {{ t('home.buyPlans') }}
                <Icon name="arrowRight" size="md" class="ml-2" :stroke-width="2" />
              </router-link>
              <router-link
                :to="isAuthenticated ? dashboardPath : '/login'"
                class="hero-link-card inline-flex items-center justify-center gap-2 rounded-xl border border-white/10 bg-dark-800/70 px-5 py-3 font-mono text-sm text-primary-300 transition-colors hover:border-primary-500/30 hover:bg-dark-800"
              >
                {{ isAuthenticated ? t('home.goToDashboard') : 'vip.aipro.love' }}
                <Icon name="arrowRight" size="sm" :stroke-width="2" />
              </router-link>
            </div>
          </div>

          <div class="flex min-w-0 w-full max-w-full justify-center overflow-hidden">
            <div class="brand-console">
              <div class="brand-mark-panel">
                <div class="station-grid" aria-hidden="true"></div>
                <div class="station-halo station-halo-blue" aria-hidden="true"></div>
                <div class="station-halo station-halo-orange" aria-hidden="true"></div>
                <svg class="route-flow" viewBox="0 0 420 300" aria-hidden="true">
                  <defs>
                    <linearGradient id="route-blue" x1="0" y1="0" x2="1" y2="0">
                      <stop offset="0%" stop-color="#079fd3" stop-opacity="0" />
                      <stop offset="45%" stop-color="#29b8ea" />
                      <stop offset="100%" stop-color="#7adaff" />
                    </linearGradient>
                    <linearGradient id="route-orange" x1="0" y1="0" x2="1" y2="1">
                      <stop offset="0%" stop-color="#f47c20" />
                      <stop offset="100%" stop-color="#fdba74" />
                    </linearGradient>
                  </defs>
                  <path class="route-flow-track" d="M58 150 H176 C195 150 199 132 210 132 C221 132 225 150 244 150 H362" />
                  <path class="route-flow-track" d="M210 48 V112 M210 188 V252" />
                  <path class="route-flow-blue" d="M58 150 H176 C195 150 199 132 210 132 C221 132 225 150 244 150 H362" />
                  <path class="route-flow-orange" d="M210 48 V112 M210 188 V252" />
                  <circle class="route-pulse route-pulse-blue" cx="88" cy="150" r="5" />
                  <circle class="route-pulse route-pulse-orange" cx="210" cy="70" r="5" />
                </svg>
                <div class="gateway-core">
                  <img :src="siteLogo || '/brand-mark.svg'" alt="Brand mark" class="gateway-core-logo" />
                </div>
                <div class="gateway-node gateway-node-left">
                  <span>API</span>
                </div>
                <div class="gateway-node gateway-node-top">
                  <span>Key</span>
                </div>
                <div class="gateway-node gateway-node-right">
                  <span>AI</span>
                </div>
                <div class="gateway-node gateway-node-bottom">
                  <span>Usage</span>
                </div>
                <div class="provider-chip provider-chip-claude">Claude</div>
                <div class="provider-chip provider-chip-gpt">GPT</div>
                <div class="provider-chip provider-chip-gemini">Gemini</div>
              </div>
              <div class="terminal-window">
                <!-- Window header -->
                <div class="terminal-header">
                  <div class="terminal-buttons">
                    <span class="btn-close"></span>
                    <span class="btn-minimize"></span>
                    <span class="btn-maximize"></span>
                  </div>
                  <span class="terminal-title">terminal</span>
                </div>
                <!-- Terminal content -->
                <div class="terminal-body">
                  <div class="code-line line-1">
                    <span class="code-prompt">$</span>
                    <span class="code-cmd">curl</span>
                    <span class="code-flag">-X POST</span>
                    <span class="code-url">/v1/messages</span>
                  </div>
                  <div class="code-line line-2">
                    <span class="code-comment"># Routing to upstream...</span>
                  </div>
                  <div class="code-line line-3">
                    <span class="code-success">200 OK</span>
                    <span class="code-response">{ "content": "Hello!" }</span>
                  </div>
                  <div class="code-line line-4">
                    <span class="code-prompt">$</span>
                    <span class="cursor"></span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Feature Tags - Centered -->
        <div class="mb-12 flex min-w-0 flex-wrap items-center justify-center gap-3 md:gap-6">
          <div
            class="inline-flex max-w-full items-center gap-2.5 rounded-lg border border-dark-100 bg-white/90 px-3 py-2.5 shadow-sm backdrop-blur-sm dark:border-dark-700/70 dark:bg-dark-900/80 sm:px-5"
          >
            <Icon name="swap" size="sm" class="text-primary-500" />
            <span class="min-w-0 text-sm font-medium text-gray-700 dark:text-dark-200">{{
              t('home.tags.subscriptionToApi')
            }}</span>
          </div>
          <div
            class="inline-flex max-w-full items-center gap-2.5 rounded-lg border border-dark-100 bg-white/90 px-3 py-2.5 shadow-sm backdrop-blur-sm dark:border-dark-700/70 dark:bg-dark-900/80 sm:px-5"
          >
            <Icon name="shield" size="sm" class="text-primary-500" />
            <span class="min-w-0 text-sm font-medium text-gray-700 dark:text-dark-200">{{
              t('home.tags.stickySession')
            }}</span>
          </div>
          <div
            class="inline-flex max-w-full items-center gap-2.5 rounded-lg border border-dark-100 bg-white/90 px-3 py-2.5 shadow-sm backdrop-blur-sm dark:border-dark-700/70 dark:bg-dark-900/80 sm:px-5"
          >
            <Icon name="chart" size="sm" class="text-primary-500" />
            <span class="min-w-0 text-sm font-medium text-gray-700 dark:text-dark-200">{{
              t('home.tags.realtimeBilling')
            }}</span>
          </div>
        </div>

        <!-- Gateway Workflow -->
        <section class="mb-12 grid min-w-0 gap-6 overflow-hidden lg:grid-cols-[0.92fr_1.08fr]">
          <div class="min-w-0 rounded-lg border border-dark-100 bg-white/88 p-6 shadow-xl shadow-dark-900/5 backdrop-blur-sm dark:border-dark-700/70 dark:bg-dark-900/78 md:p-8">
            <h2 class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ t('home.workflow.title') }}
            </h2>
            <p class="mt-2 text-sm leading-6 text-gray-600 dark:text-dark-400">
              {{ t('home.workflow.description') }}
            </p>
            <div class="mt-7 space-y-4">
              <div
                v-for="(step, index) in gatewaySteps"
                :key="step.title"
                class="min-w-0 flex gap-4 rounded-lg border border-dark-100 bg-dark-50/70 p-4 dark:border-dark-700/70 dark:bg-dark-950/50"
              >
                <div class="flex h-9 w-9 shrink-0 items-center justify-center rounded-lg bg-dark-900 text-sm font-semibold text-white dark:bg-white dark:text-dark-950">
                  {{ index + 1 }}
                </div>
                <div>
                  <h3 class="text-sm font-semibold text-gray-900 dark:text-white">{{ step.title }}</h3>
                  <p class="mt-1 text-sm leading-6 text-gray-600 dark:text-dark-400">{{ step.description }}</p>
                </div>
              </div>
            </div>
          </div>

          <div class="min-w-0 overflow-hidden rounded-lg border border-dark-100 bg-[#0c151c] p-5 shadow-2xl shadow-dark-900/20 dark:border-primary-900/40 md:p-6">
            <div class="mb-4 flex items-center justify-between border-b border-white/10 pb-3">
              <div class="flex items-center gap-2">
                <span class="h-2.5 w-2.5 rounded-full bg-accent-500"></span>
                <span class="h-2.5 w-2.5 rounded-full bg-primary-500"></span>
                <span class="h-2.5 w-2.5 rounded-full bg-white/45"></span>
              </div>
              <span class="font-mono text-xs text-white/45">{{ t('home.integration.consoleTitle') }}</span>
            </div>
            <div class="grid min-w-0 gap-5 lg:grid-cols-[1fr_0.85fr]">
              <div class="min-w-0 overflow-x-auto font-mono text-xs leading-7 sm:text-sm">
                <p class="text-white/45"># {{ t('home.integration.copyHint') }}</p>
                <p><span class="text-accent-300">const</span> <span class="text-primary-300">client</span> <span class="text-white/60">=</span> <span class="text-accent-300">new</span> <span class="text-white">OpenAI</span><span class="text-white/60">({</span></p>
                <p class="pl-4"><span class="text-white/60">baseURL:</span> <span class="text-primary-200">'https://api.example.com/v1'</span><span class="text-white/60">,</span></p>
                <p class="pl-4"><span class="text-white/60">apiKey:</span> <span class="text-primary-200">'sk-rongjian-...'</span></p>
                <p><span class="text-white/60">})</span></p>
                <p class="mt-4"><span class="text-white/45">await</span> <span class="text-white">client.chat.completions.create</span><span class="text-white/60">({</span></p>
                <p class="pl-4"><span class="text-white/60">model:</span> <span class="text-primary-200">'claude-sonnet-4'</span><span class="text-white/60">,</span></p>
                <p class="pl-4"><span class="text-white/60">messages:</span> <span class="text-accent-200">[...]</span></p>
                <p><span class="text-white/60">})</span></p>
              </div>
              <div class="min-w-0 space-y-3">
                <div
                  v-for="endpoint in endpointRows"
                  :key="endpoint.path"
                  class="min-w-0 rounded-lg border border-white/10 bg-white/[0.04] px-3 py-2"
                >
                  <div class="break-all font-mono text-xs text-primary-200">{{ endpoint.path }}</div>
                  <div class="mt-1 text-xs text-white/55">{{ endpoint.label }}</div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- Features Grid -->
        <div class="mb-12 grid min-w-0 gap-6 md:grid-cols-3">
          <!-- Feature 1: Unified Gateway -->
          <div
            class="group rounded-lg border border-dark-100 bg-white/85 p-6 backdrop-blur-sm transition-all duration-300 hover:-translate-y-0.5 hover:border-primary-200 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700/70 dark:bg-dark-900/75"
          >
            <div
              class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-primary shadow-lg shadow-primary-500/25 transition-transform group-hover:scale-105"
            >
              <Icon name="server" size="lg" class="text-white" />
            </div>
            <h3 class="mb-2 text-lg font-semibold text-gray-900 dark:text-white">
              {{ t('home.features.unifiedGateway') }}
            </h3>
            <p class="text-sm leading-relaxed text-gray-600 dark:text-dark-400">
              {{ t('home.features.unifiedGatewayDesc') }}
            </p>
          </div>

          <!-- Feature 2: Account Pool -->
          <div
            class="group rounded-lg border border-dark-100 bg-white/85 p-6 backdrop-blur-sm transition-all duration-300 hover:-translate-y-0.5 hover:border-primary-200 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700/70 dark:bg-dark-900/75"
          >
            <div
              class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-accent shadow-lg shadow-accent-500/25 transition-transform group-hover:scale-105"
            >
              <svg
                class="h-6 w-6 text-white"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="1.5"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M18 18.72a9.094 9.094 0 003.741-.479 3 3 0 00-4.682-2.72m.94 3.198l.001.031c0 .225-.012.447-.037.666A11.944 11.944 0 0112 21c-2.17 0-4.207-.576-5.963-1.584A6.062 6.062 0 016 18.719m12 0a5.971 5.971 0 00-.941-3.197m0 0A5.995 5.995 0 0012 12.75a5.995 5.995 0 00-5.058 2.772m0 0a3 3 0 00-4.681 2.72 8.986 8.986 0 003.74.477m.94-3.197a5.971 5.971 0 00-.94 3.197M15 6.75a3 3 0 11-6 0 3 3 0 016 0zm6 3a2.25 2.25 0 11-4.5 0 2.25 2.25 0 014.5 0zm-13.5 0a2.25 2.25 0 11-4.5 0 2.25 2.25 0 014.5 0z"
                />
              </svg>
            </div>
            <h3 class="mb-2 text-lg font-semibold text-gray-900 dark:text-white">
              {{ t('home.features.multiAccount') }}
            </h3>
            <p class="text-sm leading-relaxed text-gray-600 dark:text-dark-400">
              {{ t('home.features.multiAccountDesc') }}
            </p>
          </div>

          <!-- Feature 3: Billing & Quota -->
          <div
            class="group rounded-lg border border-dark-100 bg-white/85 p-6 backdrop-blur-sm transition-all duration-300 hover:-translate-y-0.5 hover:border-primary-200 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700/70 dark:bg-dark-900/75"
          >
            <div
              class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-dark-800 to-dark-950 shadow-lg shadow-dark-900/25 transition-transform group-hover:scale-105"
            >
              <svg
                class="h-6 w-6 text-white"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="1.5"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M2.25 18.75a60.07 60.07 0 0115.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 013 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 00-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 01-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 003 15h-.75M15 10.5a3 3 0 11-6 0 3 3 0 016 0zm3 0h.008v.008H18V10.5zm-12 0h.008v.008H6V10.5z"
                />
              </svg>
            </div>
            <h3 class="mb-2 text-lg font-semibold text-gray-900 dark:text-white">
              {{ t('home.features.balanceQuota') }}
            </h3>
            <p class="text-sm leading-relaxed text-gray-600 dark:text-dark-400">
              {{ t('home.features.balanceQuotaDesc') }}
            </p>
          </div>
        </div>

        <!-- Model Marketplace Inspired Section -->
        <section class="mb-12 min-w-0 overflow-hidden rounded-lg border border-dark-100 bg-white/88 p-6 shadow-xl shadow-dark-900/5 backdrop-blur-sm dark:border-dark-700/70 dark:bg-dark-900/78 md:p-8">
          <div class="mb-7 flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
            <div>
              <h2 class="text-2xl font-bold text-gray-900 dark:text-white">
                {{ t('home.modelHub.title') }}
              </h2>
              <p class="mt-2 max-w-2xl text-sm leading-6 text-gray-600 dark:text-dark-400">
                {{ t('home.modelHub.description') }}
              </p>
            </div>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="capability in modelCapabilities"
                :key="capability"
                class="rounded bg-primary-50 px-2.5 py-1 text-xs font-medium text-primary-700 dark:bg-primary-900/30 dark:text-primary-300"
              >
                {{ capability }}
              </span>
            </div>
          </div>

          <div class="grid min-w-0 gap-6 lg:grid-cols-[1.1fr_0.9fr]">
            <div class="grid min-w-0 gap-3 sm:grid-cols-2">
              <div
                v-for="provider in providerHighlights"
                :key="provider.name"
                class="min-w-0 flex items-center gap-3 rounded-lg border border-dark-100 bg-dark-50/70 p-4 dark:border-dark-700/70 dark:bg-dark-950/50"
              >
                <div
                  class="flex h-11 w-11 shrink-0 items-center justify-center rounded-lg text-sm font-bold text-white"
                  :class="provider.className"
                >
                  {{ provider.initial }}
                </div>
                <div class="min-w-0">
                  <div class="font-semibold text-gray-900 dark:text-white">{{ provider.name }}</div>
                  <div class="truncate text-xs text-gray-500 dark:text-dark-400">{{ provider.models }}</div>
                </div>
              </div>
            </div>
            <div class="min-w-0 rounded-lg bg-dark-950 p-5 text-white">
              <h3 class="text-base font-semibold">{{ t('home.modelHub.routingTitle') }}</h3>
              <p class="mt-2 text-sm leading-6 text-white/60">{{ t('home.modelHub.routingDesc') }}</p>
              <div class="mt-5 space-y-3">
                <div
                  v-for="route in routePolicies"
                  :key="route.title"
                  class="flex items-center justify-between border-b border-white/10 pb-3 last:border-b-0 last:pb-0"
                >
                  <span class="text-sm text-white/72">{{ route.title }}</span>
                  <span class="font-mono text-xs text-primary-200">{{ route.value }}</span>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- Operations Section -->
        <section class="mb-12 grid min-w-0 gap-4 md:grid-cols-4">
          <article
            v-for="item in operationItems"
            :key="item.title"
            class="rounded-lg border border-dark-100 bg-white/85 p-5 backdrop-blur-sm dark:border-dark-700/70 dark:bg-dark-900/75"
          >
            <div class="mb-4 h-1.5 w-12 rounded-full" :class="item.accent"></div>
            <h3 class="text-sm font-semibold text-gray-900 dark:text-white">{{ item.title }}</h3>
            <p class="mt-2 text-sm leading-6 text-gray-600 dark:text-dark-400">{{ item.description }}</p>
          </article>
        </section>

        <!-- Subscription Plans -->
        <section
          v-if="homepagePlans.length || plansLoading"
          class="mb-16 rounded-lg border border-dark-100 bg-white/88 p-6 shadow-xl shadow-dark-900/5 backdrop-blur-sm dark:border-dark-700/70 dark:bg-dark-900/78 md:p-8"
        >
          <div class="mb-7 flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
            <div>
              <p class="mb-2 text-sm font-semibold text-primary-600 dark:text-primary-400">
                {{ t('home.pricing.eyebrow') }}
              </p>
              <h2 class="text-2xl font-bold text-gray-900 dark:text-white">
                {{ t('home.pricing.title') }}
              </h2>
              <p class="mt-2 max-w-2xl text-sm leading-6 text-gray-600 dark:text-dark-400">
                {{ t('home.pricing.description') }}
              </p>
            </div>
            <router-link
              :to="isAuthenticated ? '/purchase' : '/login'"
              class="inline-flex items-center justify-center rounded-lg bg-dark-900 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-dark-800 dark:bg-white dark:text-dark-950 dark:hover:bg-dark-100"
            >
              {{ isAuthenticated ? t('home.pricing.cta') : t('home.pricing.loginCta') }}
              <Icon name="arrowRight" size="sm" class="ml-2" :stroke-width="2" />
            </router-link>
          </div>

          <div v-if="plansLoading" class="grid gap-4 md:grid-cols-3">
            <div
              v-for="index in 3"
              :key="index"
              class="h-56 animate-pulse rounded-lg border border-dark-100 bg-dark-50/80 dark:border-dark-700 dark:bg-dark-800/70"
            ></div>
          </div>

          <div v-else class="grid gap-4 md:grid-cols-3">
            <article
              v-for="plan in homepagePlans"
              :key="plan.id"
              class="flex min-h-[250px] flex-col rounded-lg border border-dark-100 bg-white p-5 shadow-sm transition-all duration-300 hover:-translate-y-0.5 hover:border-primary-200 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700 dark:bg-dark-950/60"
            >
              <div class="mb-4 flex items-start justify-between gap-3">
                <div>
                  <h3 class="text-lg font-semibold text-gray-900 dark:text-white">{{ plan.name }}</h3>
                  <p class="mt-1 line-clamp-2 text-sm leading-5 text-gray-500 dark:text-dark-400">
                    {{ plan.description || t('home.pricing.defaultPlanDescription') }}
                  </p>
                </div>
                <span
                  v-if="plan.group_name"
                  class="shrink-0 rounded bg-primary-50 px-2 py-1 text-xs font-medium text-primary-700 dark:bg-primary-900/30 dark:text-primary-300"
                >
                  {{ plan.group_name }}
                </span>
              </div>

              <div class="mb-5">
                <span class="text-3xl font-bold text-gray-900 dark:text-white">¥{{ formatPrice(plan.price) }}</span>
                <span class="ml-1 text-sm text-gray-500 dark:text-dark-400">/ {{ formatValidity(plan) }}</span>
                <span
                  v-if="plan.original_price && plan.original_price > plan.price"
                  class="ml-2 text-sm text-gray-400 line-through dark:text-dark-500"
                >
                  ¥{{ formatPrice(plan.original_price) }}
                </span>
              </div>

              <div class="mb-5 grid gap-2 text-sm">
                <div
                  v-if="isQuotaPack(plan)"
                  class="flex items-center justify-between rounded-lg bg-dark-50 px-3 py-2 dark:bg-dark-900"
                >
                  <span class="text-gray-500 dark:text-dark-400">{{ t('payment.planCard.totalQuota') }}</span>
                  <span class="font-medium text-gray-900 dark:text-white">{{ formatPlanQuota(plan) }}</span>
                </div>
                <div class="flex items-center justify-between rounded-lg bg-dark-50 px-3 py-2 dark:bg-dark-900">
                  <span class="text-gray-500 dark:text-dark-400">{{ t('home.pricing.windowQuota') }}</span>
                  <span class="font-medium text-gray-900 dark:text-white">{{ formatPlanWindow(plan) }}</span>
                </div>
                <div
                  v-if="plan.rate_multiplier"
                  class="flex items-center justify-between rounded-lg bg-dark-50 px-3 py-2 dark:bg-dark-900"
                >
                  <span class="text-gray-500 dark:text-dark-400">{{ t('payment.planCard.rate') }}</span>
                  <span class="font-medium text-gray-900 dark:text-white">{{ plan.rate_multiplier }}x</span>
                </div>
              </div>

              <ul class="mt-auto space-y-2">
                <li
                  v-for="feature in visiblePlanFeatures(plan)"
                  :key="feature"
                  class="flex items-start gap-2 text-sm text-gray-600 dark:text-dark-300"
                >
                  <span class="mt-1 h-1.5 w-1.5 shrink-0 rounded-full bg-accent-500"></span>
                  <span>{{ feature }}</span>
                </li>
              </ul>
            </article>
          </div>
        </section>

        <!-- Supported Providers -->
        <div class="mb-8 text-center">
          <h2 class="mb-3 text-2xl font-bold text-gray-900 dark:text-white">
            {{ t('home.providers.title') }}
          </h2>
          <p class="text-sm text-gray-600 dark:text-dark-400">
            {{ t('home.providers.description') }}
          </p>
        </div>

        <div class="mb-16 flex flex-wrap items-center justify-center gap-4">
          <!-- Claude - Supported -->
          <div
            class="flex items-center gap-2 rounded-lg border border-primary-200 bg-white/85 px-5 py-3 ring-1 ring-primary-500/15 backdrop-blur-sm dark:border-primary-800 dark:bg-dark-900/75"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-accent"
            >
              <span class="text-xs font-bold text-white">C</span>
            </div>
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{ t('home.providers.claude') }}</span>
            <span
              class="rounded bg-primary-100 px-1.5 py-0.5 text-[10px] font-medium text-primary-600 dark:bg-primary-900/30 dark:text-primary-400"
              >{{ t('home.providers.supported') }}</span
            >
          </div>
          <!-- GPT - Supported -->
          <div
            class="flex items-center gap-2 rounded-lg border border-primary-200 bg-white/85 px-5 py-3 ring-1 ring-primary-500/15 backdrop-blur-sm dark:border-primary-800 dark:bg-dark-900/75"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-primary"
            >
              <span class="text-xs font-bold text-white">G</span>
            </div>
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">GPT</span>
            <span
              class="rounded bg-primary-100 px-1.5 py-0.5 text-[10px] font-medium text-primary-600 dark:bg-primary-900/30 dark:text-primary-400"
              >{{ t('home.providers.supported') }}</span
            >
          </div>
          <!-- Gemini - Supported -->
          <div
            class="flex items-center gap-2 rounded-lg border border-primary-200 bg-white/85 px-5 py-3 ring-1 ring-primary-500/15 backdrop-blur-sm dark:border-primary-800 dark:bg-dark-900/75"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-primary"
            >
              <span class="text-xs font-bold text-white">G</span>
            </div>
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{ t('home.providers.gemini') }}</span>
            <span
              class="rounded bg-primary-100 px-1.5 py-0.5 text-[10px] font-medium text-primary-600 dark:bg-primary-900/30 dark:text-primary-400"
              >{{ t('home.providers.supported') }}</span
            >
          </div>
          <!-- Antigravity - Supported -->
          <div
            class="flex items-center gap-2 rounded-lg border border-primary-200 bg-white/85 px-5 py-3 ring-1 ring-primary-500/15 backdrop-blur-sm dark:border-primary-800 dark:bg-dark-900/75"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-accent"
            >
              <span class="text-xs font-bold text-white">A</span>
            </div>
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{ t('home.providers.antigravity') }}</span>
            <span
              class="rounded bg-primary-100 px-1.5 py-0.5 text-[10px] font-medium text-primary-600 dark:bg-primary-900/30 dark:text-primary-400"
              >{{ t('home.providers.supported') }}</span
            >
          </div>
          <!-- More - Coming Soon -->
          <div
            class="flex items-center gap-2 rounded-lg border border-dark-100 bg-white/60 px-5 py-3 opacity-70 backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-900/50"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-gray-500 to-gray-600"
            >
              <span class="text-xs font-bold text-white">+</span>
            </div>
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{ t('home.providers.more') }}</span>
            <span
              class="rounded bg-gray-100 px-1.5 py-0.5 text-[10px] font-medium text-gray-500 dark:bg-dark-700 dark:text-dark-400"
              >{{ t('home.providers.soon') }}</span
            >
          </div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <footer class="relative z-10 border-t border-gray-200/50 px-4 py-8 dark:border-dark-800/50 sm:px-6">
      <div
        class="mx-auto flex max-w-6xl flex-col items-center justify-center gap-4 text-center sm:flex-row sm:text-left"
      >
        <p class="text-sm text-gray-500 dark:text-dark-400">
          &copy; {{ currentYear }} {{ brandName }}. {{ t('home.footer.allRightsReserved') }}
        </p>
        <div class="flex items-center gap-4">
          <router-link
            to="/legal/terms"
            class="text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-dark-400 dark:hover:text-white"
          >
            {{ t('home.terms') }}
          </router-link>
          <router-link
            to="/legal/privacy"
            class="text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-dark-400 dark:hover:text-white"
          >
            {{ t('home.privacy') }}
          </router-link>
          <a
            v-if="docUrl"
            :href="docUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-dark-400 dark:hover:text-white"
          >
            {{ t('home.docs') }}
          </a>
          <a
            :href="githubUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-dark-400 dark:hover:text-white"
          >
            GitHub
          </a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import DOMPurify from 'dompurify'
import { useAuthStore, useAppStore } from '@/stores'
import { paymentAPI } from '@/api/payment'
import type { SubscriptionPlan } from '@/types/payment'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import Icon from '@/components/icons/Icon.vue'

const { t } = useI18n()

type HomepagePlan = Omit<SubscriptionPlan, 'features'> & { features: string[] }
type RawHomepagePlan = Omit<SubscriptionPlan, 'features'> & { features?: string[] | string | null }

const authStore = useAuthStore()
const appStore = useAppStore()

// Site settings - directly from appStore (already initialized from injected config)
const brandName = '熔鉴AI'
const siteLogo = computed(() => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || '')
const siteSubtitle = computed(() => appStore.cachedPublicSettings?.site_subtitle || 'AI API Gateway Platform')
const heroSubtitle = computed(() => {
  const subtitle = siteSubtitle.value.trim()
  if (!subtitle || /sub2api|subscription to api|api gateway platform/i.test(subtitle)) return t('home.heroSubtitle')
  return subtitle
})
const docUrl = computed(() => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || '')
const homeContent = computed(() => appStore.cachedPublicSettings?.home_content || '')
const sanitizedHomeContent = computed(() => DOMPurify.sanitize(homeContent.value))

// Check if homeContent is a URL (for iframe display)
const isHomeContentUrl = computed(() => {
  const content = homeContent.value.trim()
  return content.startsWith('http://') || content.startsWith('https://')
})

// Theme
const isDark = ref(document.documentElement.classList.contains('dark'))

// GitHub URL
const githubUrl = 'https://github.com/Wei-Shaw/sub2api'

// Auth state
const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => isAdmin.value ? '/admin/dashboard' : '/dashboard')
const userInitial = computed(() => {
  const user = authStore.user
  if (!user || !user.email) return ''
  return user.email.charAt(0).toUpperCase()
})

const homepagePlans = ref<HomepagePlan[]>([])
const plansLoading = ref(false)

const gatewaySteps = computed(() => [
  {
    title: t('home.workflow.steps.plan.title'),
    description: t('home.workflow.steps.plan.description')
  },
  {
    title: t('home.workflow.steps.key.title'),
    description: t('home.workflow.steps.key.description')
  },
  {
    title: t('home.workflow.steps.call.title'),
    description: t('home.workflow.steps.call.description')
  }
])

const endpointRows = computed(() => [
  { path: '/v1/chat/completions', label: t('home.integration.endpoints.chat') },
  { path: '/v1/responses', label: t('home.integration.endpoints.responses') },
  { path: '/v1/messages', label: t('home.integration.endpoints.messages') },
  { path: '/v1/images/generations', label: t('home.integration.endpoints.images') }
])

const modelCapabilities = computed(() => [
  t('home.modelHub.capabilities.chat'),
  t('home.modelHub.capabilities.code'),
  t('home.modelHub.capabilities.vision'),
  t('home.modelHub.capabilities.image'),
  t('home.modelHub.capabilities.audio')
])

const providerHighlights = computed(() => [
  {
    name: 'Claude',
    initial: 'C',
    models: t('home.modelHub.providers.claude'),
    className: 'bg-gradient-accent'
  },
  {
    name: 'GPT',
    initial: 'G',
    models: t('home.modelHub.providers.gpt'),
    className: 'bg-gradient-primary'
  },
  {
    name: 'Gemini',
    initial: 'G',
    models: t('home.modelHub.providers.gemini'),
    className: 'bg-gradient-primary'
  },
  {
    name: 'Antigravity',
    initial: 'A',
    models: t('home.modelHub.providers.antigravity'),
    className: 'bg-gradient-accent'
  }
])

const routePolicies = computed(() => [
  { title: t('home.modelHub.policies.failover'), value: 'auto' },
  { title: t('home.modelHub.policies.window'), value: 'quota' },
  { title: t('home.modelHub.policies.billing'), value: 'realtime' }
])

const operationItems = computed(() => [
  {
    title: t('home.operations.items.keys.title'),
    description: t('home.operations.items.keys.description'),
    accent: 'bg-primary-500'
  },
  {
    title: t('home.operations.items.usage.title'),
    description: t('home.operations.items.usage.description'),
    accent: 'bg-accent-500'
  },
  {
    title: t('home.operations.items.groups.title'),
    description: t('home.operations.items.groups.description'),
    accent: 'bg-dark-900 dark:bg-white'
  },
  {
    title: t('home.operations.items.store.title'),
    description: t('home.operations.items.store.description'),
    accent: 'bg-primary-500'
  }
])

// Current year for footer
const currentYear = computed(() => new Date().getFullYear())

// Toggle theme
function toggleTheme() {
  isDark.value = !isDark.value
  document.documentElement.classList.toggle('dark', isDark.value)
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
}

// Initialize theme
function initTheme() {
  const savedTheme = localStorage.getItem('theme')
  if (
    savedTheme === 'dark' ||
    (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)
  ) {
    isDark.value = true
    document.documentElement.classList.add('dark')
  }
}

function normalizePlanFeatures(features: RawHomepagePlan['features']): string[] {
  if (Array.isArray(features)) return features.filter(Boolean)
  if (!features) return []

  try {
    const parsed = JSON.parse(features)
    if (Array.isArray(parsed)) return parsed.filter(Boolean)
  } catch {
    // Plain newline-separated feature lists are also accepted by older APIs.
  }

  return features
    .split('\n')
    .map(item => item.trim())
    .filter(Boolean)
}

function normalizeHomepagePlan(plan: RawHomepagePlan): HomepagePlan {
  return {
    ...plan,
    features: normalizePlanFeatures(plan.features)
  }
}

async function loadHomepagePlans() {
  plansLoading.value = true
  try {
    const { data } = await paymentAPI.getPublicPlans()
    homepagePlans.value = ((data || []) as RawHomepagePlan[])
      .filter(plan => plan.for_sale !== false)
      .map(normalizeHomepagePlan)
      .sort((left, right) => (left.sort_order || 0) - (right.sort_order || 0) || left.id - right.id)
  } catch {
    homepagePlans.value = []
  } finally {
    plansLoading.value = false
  }
}

function formatPrice(value: number) {
  return Number(value || 0).toLocaleString(undefined, {
    minimumFractionDigits: Number.isInteger(value) ? 0 : 2,
    maximumFractionDigits: 2
  })
}

function formatValidity(plan: HomepagePlan) {
  if (isQuotaPack(plan)) return formatPlanQuota(plan)
  const days = Number(plan.validity_days || 0)
  if (days <= 0) return t('payment.planCard.unlimited')
  if (days === 30) return t('payment.oneMonth')
  if (days === 365) return t('payment.oneYear')
  return `${days} ${t('payment.days')}`
}

function isQuotaPack(plan: HomepagePlan) {
  return plan.plan_type === 'quota_pack'
}

function formatPlanQuota(plan: HomepagePlan) {
  return t('payment.planCard.requests', { count: Number(plan.quota_count || 0) })
}

function formatPlanWindow(plan: HomepagePlan) {
  const count = Number(plan.window_quota_count || 0)
  const minutes = Number(plan.window_quota_minutes || 0)
  if (count <= 0 || minutes <= 0) return t('payment.planCard.unlimited')

  const windowText = minutes % 60 === 0
    ? t('payment.planCard.hours', { count: minutes / 60 })
    : t('payment.planCard.minutes', { count: minutes })

  return t('payment.planCard.windowQuotaValue', { count, window: windowText })
}

function visiblePlanFeatures(plan: HomepagePlan) {
  if (plan.features.length) return plan.features.slice(0, 4)
  return [
    t('home.pricing.featureApiAccess'),
    t('home.pricing.featureWindowQuota'),
    t('home.pricing.featureUnifiedKey')
  ]
}

onMounted(() => {
  initTheme()

  // Check auth state
  authStore.checkAuth()

  loadHomepagePlans()

  // Ensure public settings are loaded (will use cache if already loaded from injected config)
  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings()
  }
})
</script>

<style scoped>
.brand-home {
  --brand-blue: #079fd3;
  --brand-orange: #f47c20;
  --brand-graphite: #0c151c;
}

.hero-gradient-text {
  background: linear-gradient(135deg, #ffffff 0%, #7adaff 42%, #fdba74 100%);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  -webkit-text-fill-color: transparent;
  text-shadow: 0 22px 60px rgba(7, 159, 211, 0.16);
}

.hero-link-card {
  box-shadow:
    0 16px 42px rgba(0, 0, 0, 0.16),
    inset 0 1px 0 rgba(255, 255, 255, 0.06);
}

.brand-diagonal {
  position: absolute;
  height: 2px;
  background: linear-gradient(90deg, transparent, rgba(7, 159, 211, 0.22), rgba(244, 124, 32, 0.2), transparent);
  transform: rotate(-18deg);
}

.brand-diagonal-top {
  left: 8%;
  top: 18%;
  width: 42rem;
}

.brand-diagonal-bottom {
  bottom: 14%;
  right: 0;
  width: 34rem;
}

.brand-console {
  position: relative;
  display: grid;
  min-height: 430px;
  width: min(100%, 560px);
  align-items: end;
  justify-items: center;
}

.brand-mark-panel {
  position: absolute;
  right: 0;
  top: 0;
  width: min(520px, 100%);
  height: 330px;
  border: 1px solid rgba(122, 218, 255, 0.16);
  border-radius: 8px;
  background:
    radial-gradient(circle at 50% 48%, rgba(122, 218, 255, 0.22), transparent 28%),
    radial-gradient(circle at 74% 22%, rgba(244, 124, 32, 0.14), transparent 22%),
    linear-gradient(145deg, rgba(31, 45, 54, 0.86), rgba(12, 21, 28, 0.96));
  box-shadow:
    0 36px 90px rgba(0, 0, 0, 0.24),
    inset 0 1px 0 rgba(255, 255, 255, 0.08);
  overflow: hidden;
}

.station-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(122, 218, 255, 0.07) 1px, transparent 1px),
    linear-gradient(90deg, rgba(122, 218, 255, 0.07) 1px, transparent 1px);
  background-size: 38px 38px;
  mask-image: radial-gradient(circle at center, #000 0 62%, transparent 82%);
}

.station-halo {
  position: absolute;
  inset: 50%;
  width: 250px;
  height: 250px;
  border: 1px solid;
  border-radius: 999px;
  transform: translate(-50%, -50%);
}

.station-halo-blue {
  border-color: rgba(41, 184, 234, 0.3);
  box-shadow: 0 0 42px rgba(7, 159, 211, 0.14);
}

.station-halo-orange {
  width: 178px;
  height: 178px;
  border-color: rgba(244, 124, 32, 0.32);
  transform: translate(-50%, -50%) rotate(22deg);
  border-style: dashed;
}

.route-flow {
  position: absolute;
  inset: 14px 18px;
  fill: none;
}

.route-flow-track {
  stroke: rgba(255, 255, 255, 0.1);
  stroke-linecap: square;
  stroke-width: 18;
}

.route-flow-blue {
  stroke: url(#route-blue);
  stroke-linecap: square;
  stroke-width: 8;
  stroke-dasharray: 220 48;
  animation: route-drift 5s linear infinite;
}

.route-flow-orange {
  stroke: url(#route-orange);
  stroke-linecap: square;
  stroke-width: 8;
  stroke-dasharray: 84 34;
  animation: route-drift 3.8s linear infinite;
}

.route-pulse {
  filter: drop-shadow(0 0 10px currentColor);
  animation: pulse-node 2.2s ease-in-out infinite;
}

.route-pulse-blue {
  color: #7adaff;
  fill: #7adaff;
}

.route-pulse-orange {
  color: #fdba74;
  fill: #fdba74;
  animation-delay: 0.5s;
}

.gateway-core {
  position: absolute;
  left: 50%;
  top: 50%;
  display: grid;
  height: 94px;
  width: 94px;
  place-items: center;
  transform: translate(-50%, -50%);
  border: 1px solid rgba(122, 218, 255, 0.34);
  border-radius: 8px;
  background: linear-gradient(145deg, rgba(255, 255, 255, 0.96), rgba(222, 238, 248, 0.94));
  box-shadow:
    0 18px 44px rgba(0, 0, 0, 0.28),
    0 0 0 8px rgba(7, 159, 211, 0.08);
}

.gateway-core-logo {
  height: 76px;
  width: 76px;
  object-fit: contain;
}

.gateway-node {
  position: absolute;
  display: grid;
  min-width: 72px;
  height: 38px;
  place-items: center;
  border: 1px solid rgba(122, 218, 255, 0.22);
  border-radius: 8px;
  background: rgba(12, 21, 28, 0.74);
  color: rgba(255, 255, 255, 0.8);
  font-family: ui-monospace, monospace;
  font-size: 12px;
  box-shadow: 0 12px 26px rgba(0, 0, 0, 0.18);
}

.gateway-node-left {
  left: 34px;
  top: 132px;
}

.gateway-node-top {
  left: 50%;
  top: 26px;
  transform: translateX(-50%);
}

.gateway-node-right {
  right: 34px;
  top: 132px;
}

.gateway-node-bottom {
  left: 50%;
  bottom: 26px;
  transform: translateX(-50%);
}

.provider-chip {
  position: absolute;
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.08);
  padding: 7px 10px;
  color: rgba(255, 255, 255, 0.72);
  font-size: 12px;
  line-height: 1;
  backdrop-filter: blur(10px);
}

.provider-chip-claude {
  left: 76px;
  bottom: 58px;
}

.provider-chip-gpt {
  right: 92px;
  bottom: 58px;
}

.provider-chip-gemini {
  right: 74px;
  top: 70px;
}

@keyframes route-drift {
  to {
    stroke-dashoffset: -268;
  }
}

@keyframes pulse-node {
  0%,
  100% {
    opacity: 0.45;
    transform: scale(0.8);
  }
  50% {
    opacity: 1;
    transform: scale(1.2);
  }
}

/* Terminal Window */
.terminal-window {
  position: relative;
  z-index: 2;
  width: 460px;
  max-width: min(460px, 92vw);
  background: linear-gradient(145deg, #1f2d36 0%, #0c151c 100%);
  border-radius: 8px;
  box-shadow:
    0 25px 50px -12px rgba(12, 21, 28, 0.45),
    0 0 0 1px rgba(255, 255, 255, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  overflow: hidden;
  transform: translate(-24px, 34px);
  transition: transform 0.3s ease;
}

.terminal-window:hover {
  transform: translate(-24px, 26px);
}

/* Terminal Header */
.terminal-header {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  background: rgba(12, 21, 28, 0.9);
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}

.terminal-buttons {
  display: flex;
  gap: 8px;
}

.terminal-buttons span {
  width: 12px;
  height: 12px;
  border-radius: 50%;
}

.btn-close {
  background: #f47c20;
}
.btn-minimize {
  background: #079fd3;
}
.btn-maximize {
  background: #d9e0e6;
}

.terminal-title {
  flex: 1;
  text-align: center;
  font-size: 12px;
  font-family: ui-monospace, monospace;
  color: #64748b;
  margin-right: 52px;
}

/* Terminal Body */
.terminal-body {
  padding: 20px 24px;
  font-family: ui-monospace, 'Fira Code', monospace;
  font-size: 14px;
  line-height: 2;
}

.code-line {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
  opacity: 0;
  animation: line-appear 0.5s ease forwards;
}

.line-1 {
  animation-delay: 0.3s;
}
.line-2 {
  animation-delay: 1s;
}
.line-3 {
  animation-delay: 1.8s;
}
.line-4 {
  animation-delay: 2.5s;
}

@keyframes line-appear {
  from {
    opacity: 0;
    transform: translateY(5px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.code-prompt {
  color: #f47c20;
  font-weight: bold;
}
.code-cmd {
  color: #29b8ea;
}
.code-flag {
  color: #fdba74;
}
.code-url {
  color: #7adaff;
}
.code-comment {
  color: #64748b;
  font-style: italic;
}
.code-success {
  color: #7adaff;
  background: rgba(7, 159, 211, 0.16);
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 600;
}
.code-response {
  color: #fdba74;
}

/* Blinking Cursor */
.cursor {
  display: inline-block;
  width: 8px;
  height: 16px;
  background: #f47c20;
  animation: blink 1s step-end infinite;
}

@keyframes blink {
  0%,
  50% {
    opacity: 1;
  }
  51%,
  100% {
    opacity: 0;
  }
}

/* Dark mode adjustments */
:deep(.dark) .terminal-window {
  box-shadow:
    0 25px 50px -12px rgba(0, 0, 0, 0.6),
    0 0 0 1px rgba(7, 159, 211, 0.2),
    0 0 40px rgba(244, 124, 32, 0.08),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
}

:deep(.dark) .brand-mark-panel {
  border-color: rgba(122, 218, 255, 0.16);
}

@media (max-width: 640px) {
  .brand-home {
    max-width: 100vw;
    overflow-x: hidden;
  }

  .brand-console {
    width: min(100%, 358px);
    max-width: calc(100vw - 32px);
    min-height: 420px;
    justify-items: center;
    overflow: hidden;
  }

  .brand-mark-panel {
    right: auto;
    left: 50%;
    width: 100%;
    height: 280px;
    transform: translateX(-50%);
  }

  .station-halo {
    height: 206px;
    width: 206px;
  }

  .station-halo-orange {
    height: 150px;
    width: 150px;
  }

  .gateway-core {
    height: 82px;
    width: 82px;
  }

  .gateway-core-logo {
    height: 66px;
    width: 66px;
  }

  .gateway-node {
    min-width: 58px;
    height: 34px;
    font-size: 11px;
  }

  .gateway-node-left {
    left: 18px;
    top: 122px;
  }

  .gateway-node-top {
    top: 20px;
  }

  .gateway-node-right {
    right: 18px;
    top: 122px;
  }

  .gateway-node-bottom {
    bottom: 20px;
  }

  .provider-chip {
    display: none;
  }

  .terminal-window {
    width: 100%;
    max-width: 100%;
    min-width: 0;
    transform: translateY(28px);
  }

  .terminal-window:hover {
    transform: translateY(22px);
  }

  .terminal-body {
    padding: 18px 16px;
    font-size: 12px;
  }

  .code-line {
    gap: 6px;
  }
}
</style>
