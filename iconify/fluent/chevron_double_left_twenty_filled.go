package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleLeftTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.269 15.794a.75.75 0 0 1-1.06-.026l-5.002-5.25a.75.75 0 0 1 0-1.035l5.001-5.25a.75.75 0 1 1 1.086 1.034l-4.508 4.734l4.508 4.733a.75.75 0 0 1-.025 1.06Zm4 .001a.75.75 0 0 1-1.06-.026l-5.001-5.25a.75.75 0 0 1 0-1.035l5.001-5.25a.75.75 0 1 1 1.086 1.034l-4.508 4.733l4.508 4.734a.75.75 0 0 1-.025 1.06Z"/>`),
		g.Group(children),
	)
}