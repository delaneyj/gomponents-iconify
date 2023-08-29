package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlChopsticksTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.376 2.1a.75.75 0 0 1 1.024.276L11.365 11h2.27L9.1 3.124a.75.75 0 0 1 1.3-.748L15.365 11h5.933c.388 0 .702.315.702.702V12a9.98 9.98 0 0 1-.458 3H2.458A9.996 9.996 0 0 1 2 12v-.298c0-.387.314-.702.702-.702h6.933L5.1 3.124A.75.75 0 0 1 5.376 2.1ZM3.067 16.5A10 10 0 0 0 12 22a10 10 0 0 0 8.933-5.5H3.067Z"/>`),
		g.Group(children),
	)
}