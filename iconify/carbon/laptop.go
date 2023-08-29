package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 24.005H6a2.002 2.002 0 0 1-2-2v-14a2.002 2.002 0 0 1 2-2h20a2.002 2.002 0 0 1 2 2v14a2.003 2.003 0 0 1-2 2Zm-20-16v14h20v-14Zm-4 18h28v2H2z"/>`),
		g.Group(children),
	)
}