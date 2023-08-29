package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 3.5a2.5 2.5 0 1 1 5 0v5.45a3.5 3.5 0 1 1-5 0V3.5ZM7.5 2A1.5 1.5 0 0 0 6 3.5v5.887l-.166.15a2.5 2.5 0 1 0 3.333 0L9 9.386V3.5A1.5 1.5 0 0 0 7.5 2ZM8 6a.5.5 0 0 0-1 0v4.085a1.5 1.5 0 1 0 1 0V6Z"/>`),
		g.Group(children),
	)
}