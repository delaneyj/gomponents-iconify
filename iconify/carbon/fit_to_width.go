package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FitToWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m22 11l-1.41 1.41L23.17 15H8.83l2.58-2.59L10 11l-5 5l5 5l1.41-1.41L8.83 17h14.34l-2.58 2.59L22 21l5-5l-5-5z"/><path fill="currentColor" d="M28 30H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2ZM4 4v24h24V4Z"/>`),
		g.Group(children),
	)
}