package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 3.5a3 3 0 1 1 6 0v5.354a4 4 0 1 1-6 0V3.5Zm3-1.5A1.5 1.5 0 0 0 6 3.5v5.975l-.23.22a2.5 2.5 0 1 0 3.461 0L9 9.476V3.5A1.5 1.5 0 0 0 7.5 2ZM8 6a.5.5 0 0 0-1 0v4.085a1.5 1.5 0 1 0 1 0V6Z"/>`),
		g.Group(children),
	)
}