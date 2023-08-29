package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunctionFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h8v8H3V3Zm0 10h8v8H3v-8ZM13 3h8v8h-8V3Zm0 10h8v8h-8v-8Z"/>`),
		g.Group(children),
	)
}