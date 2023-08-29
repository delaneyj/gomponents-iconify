package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveRoad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 13V4h2v9h-2ZM4 20V4h2v16H4Zm7-12V4h2v4h-2Zm0 6v-4h2v4h-2Zm0 6v-4h2v4h-2Zm4.425.575l2.125-2.125l-2.125-2.1l1.425-1.425l2.125 2.125l2.125-2.125l1.4 1.425l-2.125 2.125l2.1 2.125L21.1 22l-2.15-2.125L16.825 22l-1.4-1.425Z"/>`),
		g.Group(children),
	)
}