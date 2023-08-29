package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dvi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 16a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v7.802c0 .132-.013.263-.039.392l-1.64 8.198A2 2 0 0 1 40.362 34H7.64a2 2 0 0 1-1.962-1.608l-1.64-8.198A2 2 0 0 1 4 23.802V16Zm6 8h6m5-3h2m-2 6h2m5-6h2m-2 6h2m5-6h2m-2 6h2"/>`),
		g.Group(children),
	)
}