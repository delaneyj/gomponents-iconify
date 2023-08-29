package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandZapier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h6m12 0h-6m-3-9v6m0 6v6M5.636 5.636l4.243 4.243m8.485 8.485l-4.243-4.243m4.243-8.485l-4.243 4.243m-4.242 4.242l-4.243 4.243"/>`),
		g.Group(children),
	)
}