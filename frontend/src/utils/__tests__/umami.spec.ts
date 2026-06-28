import { beforeEach, describe, expect, it, vi } from 'vitest'

async function loadUmami() {
  vi.resetModules()
  return import('@/utils/umami')
}

describe('umami tracking', () => {
  beforeEach(() => {
    document.head.innerHTML = ''
    document.body.innerHTML = ''
    vi.unstubAllEnvs()
    delete window.umami
  })

  it('does nothing when website id is not configured', async () => {
    const { initUmami, isUmamiEnabled } = await loadUmami()

    initUmami()

    expect(isUmamiEnabled()).toBe(false)
    expect(document.querySelector('script[data-umami-sub2api]')).toBeNull()
  })

  it('loads the configured umami script once', async () => {
    vi.stubEnv('VITE_UMAMI_WEBSITE_ID', 'site-123')
    vi.stubEnv('VITE_UMAMI_SCRIPT_URL', 'https://umami.example.com/script.js')
    const { initUmami } = await loadUmami()

    initUmami()
    initUmami()

    const scripts = document.querySelectorAll('script[data-umami-sub2api]')
    expect(scripts).toHaveLength(1)
    expect(scripts[0]?.getAttribute('src')).toBe('https://umami.example.com/script.js')
    expect((scripts[0] as HTMLScriptElement).dataset.websiteId).toBe('site-123')
  })

  it('tracks route page views and clickable controls', async () => {
    vi.stubEnv('VITE_UMAMI_WEBSITE_ID', 'site-123')
    const track = vi.fn()
    window.umami = { track }
    const { enableUmamiClickTracking, trackPageView } = await loadUmami()

    trackPageView('/dashboard')
    document.body.innerHTML = '<button aria-label="Create key">ignored text</button>'
    enableUmamiClickTracking()
    document.querySelector('button')?.dispatchEvent(new MouseEvent('click', { bubbles: true }))

    expect(track).toHaveBeenNthCalledWith(1, expect.objectContaining({
      website: 'site-123',
      url: '/dashboard',
    }))
    expect(track).toHaveBeenNthCalledWith(2, 'Click: Create key', expect.objectContaining({
      path: '/',
      tag: 'button',
    }))
  })
})
