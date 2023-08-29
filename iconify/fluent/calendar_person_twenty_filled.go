package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarPersonTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 7v2.401a3 3 0 0 0-3.727 4.609A2.5 2.5 0 0 0 11 16.5c0 .167.013.334.038.5H5.5A2.5 2.5 0 0 1 3 14.5V7h14Zm-2.5-4A2.5 2.5 0 0 1 17 5.5V6H3v-.5A2.5 2.5 0 0 1 5.5 3h9Zm3 9a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm1.5 4.5c0 1.245-1 2.5-3.5 2.5S12 17.75 12 16.5a1.5 1.5 0 0 1 1.5-1.5h4a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		g.Group(children),
	)
}