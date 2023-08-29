package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 9h-8v4h2v-2h3.85L13 23h2.16L20 11V9z"/>`),
		g.Group(children),
	)
}