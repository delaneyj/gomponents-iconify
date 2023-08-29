package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMultipleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 13.5V6H2v7.5A2.5 2.5 0 0 0 4.5 16h9a2.5 2.5 0 0 0 2.5-2.5Zm0-9A2.5 2.5 0 0 0 13.5 2h-9A2.5 2.5 0 0 0 2 4.5V5h14v-.5ZM5 17c.456.607 1.182 1 2 1h6.5a4.5 4.5 0 0 0 4.5-4.5v-7c0-.818-.393-1.544-1-2v9a3.5 3.5 0 0 1-3.5 3.5H5Z"/>`),
		g.Group(children),
	)
}