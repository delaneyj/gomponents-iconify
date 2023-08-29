package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatWarningSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a6 6 0 1 1 2.996 5.195l-2.338.78a.5.5 0 0 1-.639-.612l.712-2.491A5.975 5.975 0 0 1 2 8Zm6.75 2.75a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0Zm-.258-5.84A.5.5 0 0 0 7.5 5v3.5l.008.09A.5.5 0 0 0 8.5 8.5V5l-.008-.09Z"/>`),
		g.Group(children),
	)
}