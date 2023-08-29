package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextTFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 9.5A1.5 1.5 0 0 1 11.5 8h25A1.5 1.5 0 0 1 38 9.5v4a1.5 1.5 0 0 1-3 0V11h-9.5v26h3a1.5 1.5 0 0 1 0 3h-9a1.5 1.5 0 0 1 0-3h3V11H13v2.5a1.5 1.5 0 0 1-3 0v-4Z"/>`),
		g.Group(children),
	)
}