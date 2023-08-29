package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDownSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8Zm6.809 3.462a.5.5 0 0 1-.16-.106l-.003-.003l-2.5-2.5a.5.5 0 1 1 .708-.707L7.5 9.793V5a.5.5 0 0 1 1 0v4.793l1.646-1.647a.5.5 0 0 1 .708.708l-2.5 2.5l-.003.002a.499.499 0 0 1-.542.106Z"/>`),
		g.Group(children),
	)
}