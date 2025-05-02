module.exports = {
    content: [
      './pages/**/*.{ts,tsx}',
      './components/**/*.{ts,tsx}',
    ],
    darkMode: 'media', // or 'class'
    theme: {
        extend: {
            colors: {
            background: 'var(--background)',
            foreground: 'var(--foreground)',
            },
            fontFamily: {
            sans: ['var(--font-geist-sans)', 'sans-serif'],
            mono: ['var(--font-geist-mono)', 'monospace'],
            },
        },
    },
    plugins: [],
  }
  