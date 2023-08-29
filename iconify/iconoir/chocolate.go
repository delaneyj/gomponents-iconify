package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chocolate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 6.5c-3 0-4.5-.5-4.5-3.5H5v18h14V6.5Zm0 8.5H5m0-6h14m-7 12V3"/>`),
		g.Group(children),
	)
}