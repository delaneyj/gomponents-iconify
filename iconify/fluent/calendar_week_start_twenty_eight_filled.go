package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarWeekStartTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 25A3.25 3.25 0 0 1 3 21.75V6.25A3.25 3.25 0 0 1 6.25 3h15.5A3.25 3.25 0 0 1 25 6.25v15.5A3.25 3.25 0 0 1 21.75 25H6.25Zm2.5-5.25a.75.75 0 0 0 .743-.648L9.5 19V9l-.007-.102a.75.75 0 0 0-1.486 0L8 9v10l.007.102a.75.75 0 0 0 .743.648Z"/>`),
		g.Group(children),
	)
}