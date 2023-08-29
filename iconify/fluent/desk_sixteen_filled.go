package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 3A1.75 1.75 0 0 0 1 4.75V6h13v6.5a.5.5 0 0 0 1 0V4.75A1.75 1.75 0 0 0 13.25 3H2.75ZM8 7H1v4.25c0 .966.784 1.75 1.75 1.75h3.5A1.75 1.75 0 0 0 8 11.25V7ZM3 9a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2A.5.5 0 0 1 3 9Z"/>`),
		g.Group(children),
	)
}