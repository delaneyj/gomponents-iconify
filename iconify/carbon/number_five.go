package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 23h-6v-2h6v-4h-6V9h8v2h-6v4h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}