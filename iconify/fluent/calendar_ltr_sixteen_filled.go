package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarLtrSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 11.5V6H2v5.5A2.5 2.5 0 0 0 4.5 14h7a2.5 2.5 0 0 0 2.5-2.5ZM5.997 8.248a.748.748 0 1 1-1.497 0a.748.748 0 0 1 1.497 0Zm0 2.5a.748.748 0 1 1-1.497 0a.748.748 0 0 1 1.497 0Zm2.752-2.5a.748.748 0 1 1-1.497 0a.748.748 0 0 1 1.497 0Zm0 2.5a.748.748 0 1 1-1.497 0a.748.748 0 0 1 1.497 0Zm2.748-2.5a.748.748 0 1 1-1.497 0a.748.748 0 0 1 1.497 0ZM14 4.5A2.5 2.5 0 0 0 11.5 2h-7A2.5 2.5 0 0 0 2 4.5V5h12v-.5Z"/>`),
		g.Group(children),
	)
}