package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatDismissSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a6 6 0 1 1 2.996 5.195l-2.338.78a.5.5 0 0 1-.639-.612l.712-2.491A5.975 5.975 0 0 1 2 8Zm3.896-2.104l-.057.07a.5.5 0 0 0 .057.638L7.293 8L5.896 9.396l-.057.07a.5.5 0 0 0 .057.638l.07.057a.5.5 0 0 0 .638-.057L8 8.707l1.396 1.397l.07.057a.5.5 0 0 0 .638-.057l.057-.07a.5.5 0 0 0-.057-.638L8.707 8l1.397-1.396l.057-.07a.5.5 0 0 0-.057-.638l-.07-.057a.5.5 0 0 0-.638.057L8 7.293L6.604 5.896l-.07-.057a.5.5 0 0 0-.638.057Z"/>`),
		g.Group(children),
	)
}