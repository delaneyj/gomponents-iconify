package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSignalCellularOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 1l-8.31 8.31l8.31 8.3zM4.91 4.36L3.5 5.77l6.36 6.37L1 21h17.73l2 2l1.41-1.41z"/>`),
		g.Group(children),
	)
}