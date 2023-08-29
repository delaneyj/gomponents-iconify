package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderIntergender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 11.5L20 18v-4m-8.5-.5L18 20M9 4a5 5 0 1 1 0 10A5 5 0 0 1 9 4zm5 16l2-2"/>`),
		g.Group(children),
	)
}