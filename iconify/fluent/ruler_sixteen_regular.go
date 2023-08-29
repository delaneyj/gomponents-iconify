package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 3.5A1.5 1.5 0 0 0 9.5 2h-3A1.5 1.5 0 0 0 5 3.5v9A1.5 1.5 0 0 0 6.5 14h3a1.5 1.5 0 0 0 1.5-1.5v-9ZM9.5 3a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5V11h1.5a.5.5 0 0 0 0-1H6V8.5h2a.5.5 0 1 0 0-1H6V6h1.5a.5.5 0 1 0 0-1H6V3.5a.5.5 0 0 1 .5-.5h3Z"/>`),
		g.Group(children),
	)
}