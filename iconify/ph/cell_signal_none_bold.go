package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalNoneBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M52 192v8a12 12 0 0 1-24 0v-8a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}