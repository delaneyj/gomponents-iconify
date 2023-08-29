package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSignalCellularNoSim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3h-9L7.95 5.06L19 16.11zm-15.21.74L2.38 5.15L5 7.77V21h13.23l1.62 1.62l1.41-1.41z"/>`),
		g.Group(children),
	)
}