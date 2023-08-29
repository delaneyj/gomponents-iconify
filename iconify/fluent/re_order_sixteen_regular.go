package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReOrderSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 9h11a.5.5 0 0 1 .09.992L13.5 10h-11a.5.5 0 0 1-.09-.992L2.5 9h11h-11Zm0-3h11a.5.5 0 0 1 .09.992L13.5 7h-11a.5.5 0 0 1-.09-.992L2.5 6h11h-11Z"/>`),
		g.Group(children),
	)
}