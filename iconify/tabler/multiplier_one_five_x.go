package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiplierOneFiveX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16V8l-2 2m8 6h2a2 2 0 1 0 0-4h-2V8h4m-7 8v.01M17 16l4-4m0 4l-4-4"/>`),
		g.Group(children),
	)
}