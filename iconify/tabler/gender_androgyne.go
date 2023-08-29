package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderAndrogyne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 11l6-6M4 15a5 5 0 1 0 10 0a5 5 0 1 0-10 0m15-6V5h-4m1.5 5.5l-3-3"/>`),
		g.Group(children),
	)
}