package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M48 192v8a8 8 0 0 1-16 0v-8a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}