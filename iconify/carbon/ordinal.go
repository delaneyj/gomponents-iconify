package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ordinal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 26V4h-8v6h-6v6H6v10H2v2h28v-2ZM8 26v-8h4v8Zm6 0V12h4v14Zm6 0V6h4v20Z"/>`),
		g.Group(children),
	)
}