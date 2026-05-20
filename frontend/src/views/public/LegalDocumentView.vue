<template>
  <div class="min-h-screen bg-gray-50 text-gray-900 dark:bg-dark-950 dark:text-white">
    <header class="border-b border-gray-200 bg-white/95 dark:border-dark-800 dark:bg-dark-900/95">
      <div class="mx-auto flex max-w-5xl items-center justify-between gap-4 px-4 py-4 sm:px-6">
        <RouterLink to="/home" class="flex min-w-0 items-center gap-3">
          <span class="flex h-10 w-10 flex-shrink-0 items-center justify-center overflow-hidden rounded-xl bg-white shadow-sm ring-1 ring-gray-200 dark:bg-dark-800 dark:ring-dark-700">
            <img :src="siteLogo || '/brand-mark.svg'" alt="Logo" class="h-full w-full object-contain" />
          </span>
          <span class="truncate text-base font-semibold text-gray-950 dark:text-white">
            {{ siteName }}
          </span>
        </RouterLink>
        <RouterLink
          to="/login"
          class="inline-flex flex-shrink-0 items-center justify-center rounded-lg bg-primary-600 px-4 py-2 text-sm font-semibold text-white shadow-sm shadow-primary-600/20 transition hover:bg-primary-700"
        >
          登录
        </RouterLink>
      </div>
    </header>

    <main class="mx-auto max-w-4xl px-4 py-8 sm:px-6 lg:py-10">
      <div v-if="loading" class="flex min-h-[320px] items-center justify-center">
        <div class="h-8 w-8 animate-spin rounded-full border-b-2 border-primary-600"></div>
      </div>

      <section
        v-else-if="loadError"
        class="rounded-lg border border-red-200 bg-red-50 p-6 text-red-700 dark:border-red-500/30 dark:bg-red-500/10 dark:text-red-200"
      >
        <h1 class="text-lg font-semibold">文档加载失败</h1>
        <p class="mt-2 text-sm">请稍后刷新页面重试。</p>
      </section>

      <section
        v-else-if="!currentDocument"
        class="rounded-lg border border-gray-200 bg-white p-6 dark:border-dark-700 dark:bg-dark-900"
      >
        <div class="flex items-start gap-3">
          <span class="flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-md bg-gray-100 text-gray-600 dark:bg-dark-800 dark:text-dark-300">
            <Icon name="document" size="sm" />
          </span>
          <div>
            <h1 class="text-lg font-semibold text-gray-900 dark:text-white">文档不存在</h1>
            <p class="mt-2 text-sm leading-6 text-gray-600 dark:text-dark-300">
              当前条款文档不存在或已被管理员移除。
            </p>
          </div>
        </div>
      </section>

      <article v-else>
        <div class="mb-8 border-b border-gray-200 pb-6 dark:border-dark-700">
          <div class="flex items-start gap-4">
            <span class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-md bg-primary-50 text-primary-700 dark:bg-primary-500/10 dark:text-primary-300">
              <Icon :name="documentIcon" size="md" />
            </span>
            <div class="min-w-0">
              <p class="text-sm font-medium text-primary-700 dark:text-primary-300">登录条款</p>
              <h1 class="mt-2 break-words text-2xl font-bold tracking-normal text-gray-950 dark:text-white sm:text-3xl">
                {{ currentDocument.title }}
              </h1>
              <p v-if="updatedAt" class="mt-3 text-sm text-gray-500 dark:text-dark-400">
                更新日期：{{ updatedAt }}
              </p>
            </div>
          </div>
        </div>

        <div
          v-if="hasContent"
          class="legal-document-content"
          v-html="renderedHtml"
        ></div>
        <div
          v-else
          class="rounded-lg border border-dashed border-gray-300 bg-white px-6 py-14 text-center text-sm text-gray-500 dark:border-dark-700 dark:bg-dark-900 dark:text-dark-400"
        >
          暂无正文内容
        </div>
      </article>
    </main>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { marked } from 'marked'
import DOMPurify from 'dompurify'
import Icon from '@/components/icons/Icon.vue'
import { getPublicSettings } from '@/api/auth'
import { sanitizeUrl } from '@/utils/url'
import type { LoginAgreementDocument, PublicSettings } from '@/types'

type LegalDocumentIcon = 'document' | 'shield' | 'globe' | 'cog'

const route = useRoute()
const settings = ref<PublicSettings | null>(null)
const loading = ref(true)
const loadError = ref(false)

marked.setOptions({
  breaks: true,
  gfm: true,
})

const documentId = computed(() => String(route.params.documentId || ''))
const defaultTermsDocument: LoginAgreementDocument = {
  id: 'terms',
  title: '使用条款',
  content_md: `# 熔鉴AI 使用条款

欢迎使用熔鉴AI。使用本网站、控制台、API 中转服务、套餐购买、次卡额度或相关功能，即表示你已阅读、理解并同意本使用条款。

## 1. 服务说明

熔鉴AI提供 AI API 中转、模型接入、密钥分发、用量统计、套餐和额度管理等服务。平台可能根据上游服务状态、模型可用性、风控要求、维护安排或商业策略调整可用模型、接口、价格、额度和功能。

## 2. 账户与密钥安全

你应妥善保管账号、密码、API Key、访问令牌和相关凭据。因主动泄露、共享、被盗用或在公开仓库、截图、日志中暴露密钥导致的损失，由你自行承担。发现异常使用时，请及时停用密钥或联系平台处理。

## 3. 合规使用

你不得使用本服务从事违法违规、侵权、欺诈、攻击、滥用、绕过限制、批量注册、恶意爬取、生成或传播违法内容等行为。你应同时遵守所在地法律法规、平台规则以及相关上游模型服务商的使用政策。

## 4. 套餐、额度与计费

套餐、次卡、窗口额度、倍率和有效期以购买页面、订单记录或后台展示为准。调用成功后可能按规则扣减余额、次数或额度。因网络波动、上游限流、模型不可用或其他不可抗因素造成的短时不可用，平台会尽力恢复，但不承诺服务永不中断。

## 5. 内容与输出

AI 生成内容可能不准确、不完整或不适合特定用途。你应自行判断、验证并承担使用输出内容产生的责任。不得将 AI 输出直接用于医疗、法律、金融、身份认证、安全控制等高风险场景，除非你已进行充分人工审核并承担相应责任。

## 6. 服务变更、中止与终止

若发现违反本条款、异常高风险调用、恶意攻击、支付异常、绕过计费或影响平台稳定的行为，平台有权限制、暂停或终止相关账号、密钥、套餐或调用权限。平台也可因维护、升级、合规要求或上游变化调整服务。

## 7. 免责声明

在法律允许范围内，平台不对间接损失、利润损失、业务中断、数据丢失、第三方服务故障、上游模型变更或不可抗力造成的损失承担责任。你应自行评估服务是否适合你的业务场景。

## 8. 条款更新

平台可根据业务、合规或服务变化更新本条款。更新后的条款将在网站或登录流程中展示。继续使用服务即视为接受更新后的条款。

如你不同意本条款，请停止使用本服务。`
}

const defaultPrivacyDocument: LoginAgreementDocument = {
  id: 'privacy',
  title: '隐私政策',
  content_md: `# 熔鉴AI 隐私政策

熔鉴AI重视你的隐私和数据安全。本隐私政策说明我们在提供网站、控制台、API 中转、套餐购买、用量统计和相关服务时，如何收集、使用、保存和保护信息。

## 1. 我们收集的信息

为提供服务，我们可能收集账号信息、登录信息、联系方式、订单与支付状态、套餐与额度信息、API Key 元数据、调用时间、模型名称、请求状态、用量统计、设备与浏览器信息、IP 地址、日志和安全审计记录。

## 2. API 调用与内容数据

为完成模型转发、计费、排障和安全风控，我们可能处理与你的 API 调用相关的必要元数据。除非为完成请求转发、排查故障、履行法律义务或处理安全事件所必需，我们不会主动查看、出售或用于无关目的处理你的调用内容。

## 3. 信息使用目的

我们使用信息用于账号登录、身份验证、套餐开通、额度扣减、用量展示、订单处理、客服支持、服务稳定性监控、异常调用识别、安全防护、合规审计和产品改进。

## 4. 信息共享与第三方服务

为了提供 AI 模型转发、支付、登录验证、邮件通知、验证码、防滥用或基础设施服务，我们可能将必要信息提供给相关第三方服务商或上游模型服务商。我们会尽量限制共享范围，仅提供完成服务所需的信息。

## 5. Cookie 与本地存储

网站可能使用 Cookie、localStorage 或类似技术保存登录状态、语言偏好、主题设置、协议确认状态和安全相关信息。你可以通过浏览器设置清理这些数据，但部分功能可能因此无法正常使用。

## 6. 数据保存与安全

我们会在实现本政策目的所需期间保存相关信息，并采取合理的技术和管理措施保护数据安全。但互联网服务无法保证绝对安全，你也应妥善保管账号、密码和 API Key，避免在公开位置泄露。

## 7. 你的权利

在适用法律允许范围内，你可以请求访问、更正或删除与你相关的个人信息，也可以停止使用服务、删除密钥或联系平台处理账号与数据问题。某些记录可能因账务、风控、安全或法律合规要求需要保留一段时间。

## 8. 未成年人

本服务主要面向具备相应民事行为能力的用户。未成年人使用本服务应取得监护人同意，并在监护人指导下使用。

## 9. 政策更新

我们可能根据业务、技术、法律或合规要求更新本隐私政策。更新后的政策将在网站或相关页面展示。继续使用服务即表示你接受更新后的政策。

如你对隐私政策或数据处理有疑问，请通过网站提供的联系方式与我们联系。`
}

const documents = computed(() => {
  const configured = settings.value?.login_agreement_documents ?? []
  const hasTerms = configured.some((doc) => doc.id === defaultTermsDocument.id)
  const hasPrivacy = configured.some((doc) => doc.id === defaultPrivacyDocument.id)
  return [
    ...(hasTerms ? [] : [defaultTermsDocument]),
    ...(hasPrivacy ? [] : [defaultPrivacyDocument]),
    ...configured,
  ]
})
const siteName = computed(() => {
  const raw = settings.value?.site_name?.trim()
  return raw && !/^sub2api$/i.test(raw) ? raw : '熔鉴AI'
})
const siteLogo = computed(() => sanitizeUrl(settings.value?.site_logo || '', {
  allowRelative: true,
  allowDataUrl: true,
}))

const currentDocument = computed<LoginAgreementDocument | null>(() => {
  const id = documentId.value
  if (!id) {
    return null
  }
  const doc = documents.value.find((item) => item.id === id)
  if (id === defaultTermsDocument.id && (!doc || !doc.content_md?.trim())) {
    return defaultTermsDocument
  }
  if (id === defaultPrivacyDocument.id && (!doc || !doc.content_md?.trim())) {
    return defaultPrivacyDocument
  }
  return doc ?? null
})

const updatedAt = computed(() => {
  const doc = currentDocument.value
  const isDefaultLegalDocument =
    (doc?.id === defaultTermsDocument.id && doc.content_md === defaultTermsDocument.content_md) ||
    (doc?.id === defaultPrivacyDocument.id && doc.content_md === defaultPrivacyDocument.content_md)
  if (isDefaultLegalDocument) {
    return '2026-05-20'
  }
  return settings.value?.login_agreement_updated_at || ''
})

const hasContent = computed(() => Boolean(currentDocument.value?.content_md?.trim()))

const renderedHtml = computed(() => {
  const content = currentDocument.value?.content_md?.trim() || ''
  if (!content) {
    return ''
  }
  const html = marked.parse(content) as string
  return DOMPurify.sanitize(html)
})

const documentIcon = computed<LegalDocumentIcon>(() => {
  const title = currentDocument.value?.title || ''
  if (title.includes('政策') || title.includes('隐私')) {
    return 'shield'
  }
  if (title.includes('国家') || title.includes('地区')) {
    return 'globe'
  }
  if (title.includes('特定')) {
    return 'cog'
  }
  return 'document'
})

onMounted(async () => {
  loading.value = true
  loadError.value = false
  try {
    settings.value = await getPublicSettings()
  } catch {
    loadError.value = true
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.legal-document-content {
  line-height: 1.75;
  overflow-wrap: anywhere;
  color: inherit;
}

.legal-document-content :deep(h1) {
  @apply mb-4 mt-8 border-b border-gray-200 pb-3 text-3xl font-bold dark:border-dark-700;
}

.legal-document-content :deep(h2) {
  @apply mb-3 mt-7 text-2xl font-bold;
}

.legal-document-content :deep(h3) {
  @apply mb-2 mt-6 text-xl font-semibold;
}

.legal-document-content :deep(h4) {
  @apply mb-2 mt-5 text-lg font-semibold;
}

.legal-document-content :deep(p) {
  @apply mb-4 text-gray-700 dark:text-dark-200;
}

.legal-document-content :deep(a) {
  @apply text-primary-600 underline underline-offset-4 hover:text-primary-700 dark:text-primary-300 dark:hover:text-primary-200;
}

.legal-document-content :deep(ul) {
  @apply mb-4 list-disc pl-6;
}

.legal-document-content :deep(ol) {
  @apply mb-4 list-decimal pl-6;
}

.legal-document-content :deep(li) {
  @apply mb-1 text-gray-700 dark:text-dark-200;
}

.legal-document-content :deep(blockquote) {
  @apply my-5 border-l-4 border-gray-300 pl-4 text-gray-600 dark:border-dark-600 dark:text-dark-300;
}

.legal-document-content :deep(code) {
  @apply rounded bg-gray-100 px-1.5 py-0.5 font-mono text-sm dark:bg-dark-800;
}

.legal-document-content :deep(pre) {
  @apply my-5 overflow-x-auto rounded-lg bg-gray-950 p-4 text-gray-100;
}

.legal-document-content :deep(pre code) {
  @apply bg-transparent p-0 text-inherit;
}

.legal-document-content :deep(table) {
  @apply my-5 block w-full overflow-x-auto border-collapse;
}

.legal-document-content :deep(th) {
  @apply border border-gray-300 bg-gray-50 px-3 py-2 text-left font-semibold dark:border-dark-600 dark:bg-dark-800;
}

.legal-document-content :deep(td) {
  @apply border border-gray-300 px-3 py-2 dark:border-dark-600;
}

.legal-document-content :deep(img) {
  @apply my-5 h-auto max-w-full rounded-lg;
}

.legal-document-content :deep(hr) {
  @apply my-7 border-gray-200 dark:border-dark-700;
}
</style>
