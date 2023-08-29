package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundSignalCellularNoSim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 5c0-1.1-.9-2-2-2h-6.17c-.53 0-1.04.21-1.42.59L7.95 5.06L19 16.11V5zM3.09 4.44a.996.996 0 0 0 0 1.41L5 7.78V19a2 2 0 0 0 2 2h11.23l.91.91a.996.996 0 1 0 1.41-1.41L4.5 4.44a.996.996 0 0 0-1.41 0z"/>`),
		g.Group(children),
	)
}