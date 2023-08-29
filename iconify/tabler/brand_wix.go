package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandWix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 9l1.5 6l1.379-5.515a.64.64 0 0 1 1.242 0L8.5 15L10 9m3 2.5V15m3-6l5 6m0-6l-5 6m-3-6h.01"/>`),
		g.Group(children),
	)
}