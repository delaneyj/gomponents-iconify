package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Antox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a10.72 10.72 0 0 0-10.73 10.73v5.85h-2.92A1.94 1.94 0 0 0 8.4 23v18.55a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V23a1.94 1.94 0 0 0-2-1.94h-2.92v-5.83A10.72 10.72 0 0 0 24 4.5Zm0 2.92a7.43 7.43 0 0 1 7.8 7.81c0 4.94-3.9 8.77-8.8 12.67l.43-4.9A7.8 7.8 0 0 1 24 7.42Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.45 37.93a1.22 1.22 0 1 1 1.22-1.22a1.22 1.22 0 0 1-1.22 1.22Zm7.12 0a1.22 1.22 0 1 1 1.22-1.22a1.22 1.22 0 0 1-1.22 1.22Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 30.4h0a8.72 8.72 0 0 1 8.72 8.72v0a1.24 1.24 0 0 1-1.24 1.24h-15a1.24 1.24 0 0 1-1.24-1.24v0A8.72 8.72 0 0 1 24 30.4Zm-7.53-1.3l2.48 2.91m12.58-2.91l-2.48 2.91"/>`),
		g.Group(children),
	)
}