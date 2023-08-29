package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l9.5 5.5v11L12 23l-9.5-5.5v-11L12 1ZM4.5 7.658v8.689l7.5 4.342V12L4.5 7.658Z"/>`),
		g.Group(children),
	)
}