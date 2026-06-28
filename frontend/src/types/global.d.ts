import type { PublicSettings } from '@/types'

type UmamiTrack = {
  (eventName: string, eventData?: Record<string, unknown>): void
  (properties?: Record<string, unknown>): void
}

declare global {
  interface Window {
    __APP_CONFIG__?: PublicSettings
    umami?: {
      track?: UmamiTrack
    }
  }
}

export {}
