package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowsThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 140H48a12 12 0 0 0-12 12v40a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12v-40a12 12 0 0 0-12-12Zm4 52a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4v-40a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4Zm-4-140H48a12 12 0 0 0-12 12v40a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12Zm4 52a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4V64a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}