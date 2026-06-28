const WEBSITE_ID = import.meta.env.VITE_UMAMI_WEBSITE_ID?.trim()
const SCRIPT_URL = (import.meta.env.VITE_UMAMI_SCRIPT_URL || 'https://ottozhang.com/umami/tracker').trim()

let scriptLoaded = false
let clickTrackingEnabled = false

export function isUmamiEnabled(): boolean {
  return Boolean(WEBSITE_ID && SCRIPT_URL && typeof document !== 'undefined')
}

export function initUmami(): void {
  if (!isUmamiEnabled() || scriptLoaded || document.querySelector('script[data-website-id][data-umami-sub2api]')) return

  const script = document.createElement('script')
  script.defer = true
  script.src = SCRIPT_URL
  script.dataset.websiteId = WEBSITE_ID
  script.dataset.umamiSub2api = 'true'
  document.head.appendChild(script)
  scriptLoaded = true
}

export function trackPageView(url: string, title = document.title): void {
  if (!isUmamiEnabled()) return

  window.umami?.track?.({
    website: WEBSITE_ID,
    url,
    title,
    referrer: document.referrer,
  })
}

export function trackEvent(eventName: string, eventData?: Record<string, unknown>): void {
  if (!isUmamiEnabled()) return
  window.umami?.track?.(eventName, eventData)
}

function getClickTarget(target: EventTarget | null): HTMLElement | null {
  if (!(target instanceof Element)) return null
  return target.closest('[data-umami-event], a, button, [role="button"], input[type="button"], input[type="submit"]')
}

function getElementLabel(element: HTMLElement): string {
  const explicit = element.dataset.umamiEvent?.trim()
  if (explicit) return explicit

  const ariaLabel = element.getAttribute('aria-label')?.trim()
  if (ariaLabel) return ariaLabel

  const title = element.getAttribute('title')?.trim()
  if (title) return title

  if (element instanceof HTMLInputElement) {
    return element.value || element.name || element.type
  }

  return element.textContent?.replace(/\s+/g, ' ').trim().slice(0, 80) || element.tagName.toLowerCase()
}

function getSafeHref(element: HTMLElement): string | undefined {
  if (!(element instanceof HTMLAnchorElement) || !element.href) return undefined

  const url = new URL(element.href)
  return url.origin === window.location.origin ? url.pathname : url.origin + url.pathname
}

function handleDocumentClick(event: MouseEvent): void {
  const element = getClickTarget(event.target)
  if (!element) return

  trackEvent(`Click: ${getElementLabel(element)}`, {
    path: window.location.pathname,
    tag: element.tagName.toLowerCase(),
    href: getSafeHref(element),
  })
}

export function enableUmamiClickTracking(): void {
  if (!isUmamiEnabled() || clickTrackingEnabled) return
  document.addEventListener('click', handleDocumentClick, true)
  clickTrackingEnabled = true
}
