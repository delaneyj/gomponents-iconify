package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencySom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 18V6H5v10a2 2 0 0 1-2 2M14 6v12h4a3 3 0 0 0 0-6h-4h4a3 3 0 0 0 0-6h-4z"/>`),
		g.Group(children),
	)
}