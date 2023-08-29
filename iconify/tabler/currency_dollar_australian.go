package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyDollarAustralian(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 18L6.279 6.524a.75.75 0 0 1 1.442 0L11 18M21 6h-4a3 3 0 0 0 0 6h1a3 3 0 0 1 0 6h-4m3 2v-2m1-12V4M4.5 14h5"/>`),
		g.Group(children),
	)
}