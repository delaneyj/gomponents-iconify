package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarInfoSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.257 14H4.5A2.5 2.5 0 0 1 2 11.5V6h5.337A5.493 5.493 0 0 0 5 10.5c0 1.33.472 2.55 1.257 3.5ZM14 5v-.5A2.5 2.5 0 0 0 11.5 2h-7A2.5 2.5 0 0 0 2 4.5V5h12ZM9.875 8.5a.625.625 0 1 1 1.25 0a.625.625 0 0 1-1.25 0Zm1.125 4a.5.5 0 0 1-1 0v-2a.5.5 0 0 1 1 0v2Zm-5-2a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Zm1 0a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Z"/>`),
		g.Group(children),
	)
}