/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        // 主色调 - 品牌亮蓝
        primary: {
          50: '#effbff',
          100: '#d8f4ff',
          200: '#b8ecff',
          300: '#7adaff',
          400: '#29b8ea',
          500: '#079fd3',
          600: '#0285b5',
          700: '#036a91',
          800: '#075977',
          900: '#0b4963',
          950: '#062f42'
        },
        // 辅助色 - 品牌橙色增长线
        accent: {
          50: '#fff7ed',
          100: '#ffedd5',
          200: '#fed7aa',
          300: '#fdba74',
          400: '#fb923c',
          500: '#f47c20',
          600: '#df6511',
          700: '#b84d0f',
          800: '#933f14',
          900: '#773614',
          950: '#401907'
        },
        // 深色模式背景 - 石墨黑
        dark: {
          50: '#f7f8fa',
          100: '#eef1f4',
          200: '#d9e0e6',
          300: '#b7c4cf',
          400: '#8fa0af',
          500: '#687b89',
          600: '#4a5c68',
          700: '#344650',
          800: '#1f2d36',
          900: '#17242d',
          950: '#0c151c'
        }
      },
      fontFamily: {
        sans: [
          'system-ui',
          '-apple-system',
          'BlinkMacSystemFont',
          'Segoe UI',
          'Roboto',
          'Helvetica Neue',
          'Arial',
          'PingFang SC',
          'Hiragino Sans GB',
          'Microsoft YaHei',
          'sans-serif'
        ],
        mono: ['ui-monospace', 'SFMono-Regular', 'Menlo', 'Monaco', 'Consolas', 'monospace']
      },
      boxShadow: {
        glass: '0 8px 32px rgba(0, 0, 0, 0.08)',
        'glass-sm': '0 4px 16px rgba(0, 0, 0, 0.06)',
        glow: '0 0 20px rgba(7, 159, 211, 0.22)',
        'glow-lg': '0 0 40px rgba(244, 124, 32, 0.24)',
        card: '0 1px 3px rgba(0, 0, 0, 0.04), 0 1px 2px rgba(0, 0, 0, 0.06)',
        'card-hover': '0 10px 40px rgba(0, 0, 0, 0.08)',
        'inner-glow': 'inset 0 1px 0 rgba(255, 255, 255, 0.1)'
      },
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-primary': 'linear-gradient(135deg, #079fd3 0%, #036a91 100%)',
        'gradient-accent': 'linear-gradient(135deg, #f47c20 0%, #df6511 100%)',
        'gradient-dark': 'linear-gradient(135deg, #1f2d36 0%, #0c151c 100%)',
        'gradient-glass':
          'linear-gradient(135deg, rgba(255,255,255,0.1) 0%, rgba(255,255,255,0.05) 100%)',
        'mesh-gradient':
          'linear-gradient(135deg, rgba(7, 159, 211, 0.08) 0%, transparent 34%), linear-gradient(315deg, rgba(244, 124, 32, 0.08) 0%, transparent 30%), linear-gradient(rgba(12, 21, 28, 0.035) 1px, transparent 1px), linear-gradient(90deg, rgba(12, 21, 28, 0.035) 1px, transparent 1px)'
      },
      animation: {
        'fade-in': 'fadeIn 0.3s ease-out',
        'slide-up': 'slideUp 0.3s ease-out',
        'slide-down': 'slideDown 0.3s ease-out',
        'slide-in-right': 'slideInRight 0.3s ease-out',
        'scale-in': 'scaleIn 0.2s ease-out',
        'pulse-slow': 'pulse 3s cubic-bezier(0.4, 0, 0.6, 1) infinite',
        shimmer: 'shimmer 2s linear infinite',
        glow: 'glow 2s ease-in-out infinite alternate'
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' }
        },
        slideUp: {
          '0%': { opacity: '0', transform: 'translateY(10px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' }
        },
        slideDown: {
          '0%': { opacity: '0', transform: 'translateY(-10px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' }
        },
        slideInRight: {
          '0%': { opacity: '0', transform: 'translateX(20px)' },
          '100%': { opacity: '1', transform: 'translateX(0)' }
        },
        scaleIn: {
          '0%': { opacity: '0', transform: 'scale(0.95)' },
          '100%': { opacity: '1', transform: 'scale(1)' }
        },
        shimmer: {
          '0%': { backgroundPosition: '-200% 0' },
          '100%': { backgroundPosition: '200% 0' }
        },
        glow: {
          '0%': { boxShadow: '0 0 20px rgba(7, 159, 211, 0.22)' },
          '100%': { boxShadow: '0 0 30px rgba(244, 124, 32, 0.28)' }
        }
      },
      backdropBlur: {
        xs: '2px'
      },
      borderRadius: {
        '4xl': '2rem'
      }
    }
  },
  plugins: []
}
