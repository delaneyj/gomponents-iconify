package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 10v8v-8m1-1h-2v8h-3V9h-2v10h5v4h2v-4h1v-2h-1V9Z"/>`),
		g.Group(children),
	)
}