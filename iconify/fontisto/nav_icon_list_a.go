package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavIconListA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.216 11.998a2.608 2.608 0 1 1-5.217 0a2.608 2.608 0 0 1 5.217 0zm0-9.39a2.608 2.608 0 1 1-5.217 0a2.608 2.608 0 0 1 5.217 0zm0 18.781a2.608 2.608 0 1 1-5.217 0a2.608 2.608 0 0 1 5.217 0zM9.794 0h15.247a2.61 2.61 0 1 1 0 5.22H9.794a2.61 2.61 0 1 1 0-5.22zm0 9.39h15.247a2.61 2.61 0 1 1 0 5.22H9.794a2.61 2.61 0 1 1 0-5.22zm0 9.391h15.247a2.61 2.61 0 1 1 0 5.22H9.794a2.61 2.61 0 1 1 0-5.22z"/>`),
		g.Group(children),
	)
}