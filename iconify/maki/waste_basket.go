package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WasteBasket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m12.41 5.58l-1.34 8a.5.5 0 0 1-.49.41H4.42a.5.5 0 0 1-.49-.41l-1.34-8A.5.5 0 0 1 3.08 5h8.83a.5.5 0 0 1 .5.58zM13 3.5a.5.5 0 0 1-.5.5h-10a.5.5 0 0 1 0-1H5V1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V3h2.5a.5.5 0 0 1 .5.5zM9 3V2H6v1h3z"/>`),
		g.Group(children),
	)
}