package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 8h-1V2h-2v6h-6V2h-2v6h-1a2 2 0 0 0-2 2v6a8.007 8.007 0 0 0 7 7.93V30h2v-6.07A8.007 8.007 0 0 0 24 16v-6a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}