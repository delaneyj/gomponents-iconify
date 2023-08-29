package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toogoodtogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 9h0c2.63 0 5.15 1.8 7 5a24.65 24.65 0 0 1 2.9 12.09c0 9.43-4.43 17.08-9.9 17.08s-9.9-7.65-9.9-17.08S18.52 9 24 9Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.4 42c-3.51-.86-7.21-3.76-9.91-6.87C4.32 28 2.64 19.34 6.75 15.77a6.69 6.69 0 0 1 4.42-1.52h0a13.74 13.74 0 0 1 5.1 1.12m15.46-.01a13.74 13.74 0 0 1 5.1-1.12h0a6.69 6.69 0 0 1 4.42 1.52c4.11 3.57 2.43 12.22-3.74 19.34c-2.7 3.11-6.4 6-9.91 6.87"/>`),
		g.Group(children),
	)
}