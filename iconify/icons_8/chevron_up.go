package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 5.688l-.72.718l-13 13l-.686.688l.687.718L6.19 24.72l.718.686l.688-.687L16 16.312l8.406 8.406l.688.686l.718-.687l3.907-3.907l.686-.72l-.687-.687l-13-13l-.72-.718zm0 2.843l11.563 11.595l-2.438 2.438l-8.406-8.375L16 13.5l-.72.688l-8.405 8.374l-2.438-2.437L16 8.53z"/>`),
		g.Group(children),
	)
}