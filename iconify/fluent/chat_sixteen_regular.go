package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 7a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3A.5.5 0 0 1 6 7Zm.5 1.5a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2ZM2 8a6 6 0 1 1 2.996 5.195l-2.338.78a.5.5 0 0 1-.639-.612l.712-2.491A5.975 5.975 0 0 1 2 8Zm6-5a5 5 0 0 0-4.3 7.552a.5.5 0 0 1 .05.393l-.509 1.78l1.658-.552a.5.5 0 0 1 .426.052A5 5 0 1 0 8 3Z"/>`),
		g.Group(children),
	)
}