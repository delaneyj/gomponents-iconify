package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalNoneLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M46 192v8a6 6 0 0 1-12 0v-8a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}