package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaySixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.235 2.192A1.5 1.5 0 0 0 4 3.499v9a1.5 1.5 0 0 0 2.235 1.308l8-4.5a1.5 1.5 0 0 0 0-2.615l-8-4.5Z"/>`),
		g.Group(children),
	)
}