package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineSignalCellularThreeBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity=".3" d="M2 22h20V2z"/><path fill="currentColor" d="M17 7L2 22h15z"/>`),
		g.Group(children),
	)
}