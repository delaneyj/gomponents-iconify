package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderTrasvesti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 20a5 5 0 1 1 0-10a5 5 0 0 1 0 10zM6 6l5.4 5.4M4 8l4-4"/>`),
		g.Group(children),
	)
}