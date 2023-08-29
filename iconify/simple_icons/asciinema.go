package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asciinema(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.61 0v24l20.78-12L1.61 0m4.15 7.2l4.3 2.48l-4.3 2.48V7.2m6.79 3.92l1.53.88l-8.32 4.8v-1.76l6.79-3.92Z"/>`),
		g.Group(children),
	)
}