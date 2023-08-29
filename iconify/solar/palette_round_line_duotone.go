package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteRoundLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 6a4 4 0 1 1 8 0v12a4 4 0 0 1-8 0V6Z"/><path d="m10 8.243l3.314-3.314a4 4 0 1 1 5.657 5.657L9.306 20.25"/><path d="M6 22h12a4 4 0 0 0 0-8h-2.5M7 18a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}