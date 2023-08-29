package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ligature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 20V8c0-2.2 1.8-4 4-4c1.5 0 2.8.8 3.5 2M6 12h4m4 0h2v8M6 20h4m4 0h4"/>`),
		g.Group(children),
	)
}