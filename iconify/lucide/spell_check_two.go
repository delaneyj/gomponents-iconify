package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpellCheckTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 16l6-12l6 12M8 12h8M4 21c1.1 0 1.1-1 2.3-1s1.1 1 2.3 1c1.1 0 1.1-1 2.3-1c1.1 0 1.1 1 2.3 1c1.1 0 1.1-1 2.3-1c1.1 0 1.1 1 2.3 1c1.1 0 1.1-1 2.3-1"/>`),
		g.Group(children),
	)
}