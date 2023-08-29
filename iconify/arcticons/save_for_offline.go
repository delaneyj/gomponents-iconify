package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveForOffline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.656 6.606l4.186-.06a1.8 1.8 0 0 1 1.54.826a1.724 1.724 0 0 0 1.137.211h3.562a2.996 2.996 0 0 0 1.676-.45l12.906-.14a1.324 1.324 0 0 1 1.3.74a1.482 1.482 0 0 1 1.725-1.19q.051.01.101.023a2.172 2.172 0 0 1 1.593.666l3.235 3.183a2.414 2.414 0 0 1 .72 1.7l.21 26.489a2.32 2.32 0 0 1-.713 1.692l-.436.419a2.494 2.494 0 0 1-1.722.695l-29.569.07a2.333 2.333 0 0 1-1.691-.72l-.271-.284a2.492 2.492 0 0 1-.692-1.722V9.034a3.042 3.042 0 0 1 .527-1.751Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.963 7.734V19.06a1.004 1.004 0 0 1-1 1.004l-18.294.067a1.008 1.008 0 0 1-1.012-.996l-.138-11.552m5.561 0l.15 12.553"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.633 10.368v6.343a1 1 0 0 0 1 1h1.984a.991.991 0 0 0 .99-.99v-.01l-.057-6.343a1.01 1.01 0 0 0-1.009-1h-1.908a1 1 0 0 0-1 1ZM14.006 25.074v12.532a1 1 0 0 0 1 1H33a.995.995 0 0 0 .995-.995v-.005l-.065-12.61A1.001 1.001 0 0 0 32.924 24l-17.918.07a1.004 1.004 0 0 0-1 1.004Z"/><ellipse cx="24.341" cy="31.315" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.112" ry="5.224"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.16 27.181l-.154.308a1.891 1.891 0 0 0-.222 1.496a1.547 1.547 0 0 0-.753 2.041a1.512 1.512 0 0 0 .753.966l.602-.301l.526.376l.601 2.405m-5.902-8.05c-.069.614-.013 1.588-.582 1.83a1.937 1.937 0 0 0-.635.84c-.294.373-1.083.264-1.46.554a.746.746 0 0 0 .418.957a1.396 1.396 0 0 0 1.116.19a2.085 2.085 0 0 1 1.4.778a2.383 2.383 0 0 0 1.134.59a2.242 2.242 0 0 1-.34 2.287a5.766 5.766 0 0 1-.962 1.174l-.2.203"/>`),
		g.Group(children),
	)
}