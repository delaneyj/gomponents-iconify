package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarWeekNumbersTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.75 3A3.25 3.25 0 0 1 21 6.25v11.5A3.25 3.25 0 0 1 17.75 21H6.25A3.25 3.25 0 0 1 3 17.75V6.25A3.25 3.25 0 0 1 6.25 3h11.5ZM8.493 7.648A.75.75 0 0 0 7 7.75v1l.007.102A.75.75 0 0 0 8.5 8.75v-1l-.007-.102Zm0 7.5A.75.75 0 0 0 7 15.25v1l.007.102A.75.75 0 0 0 8.5 16.25v-1l-.007-.102Zm-.736-4.475a.75.75 0 0 0-.75.75L7 12.57l.007.102A.75.75 0 0 0 8.5 12.57l.007-1.148l-.007-.102a.75.75 0 0 0-.743-.648Z"/>`),
		g.Group(children),
	)
}