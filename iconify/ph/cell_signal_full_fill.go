package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalFullFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 40v160a16 16 0 0 1-16 16H32a16 16 0 0 1-11.3-27.32l160-160A16 16 0 0 1 208 40Z"/>`),
		g.Group(children),
	)
}