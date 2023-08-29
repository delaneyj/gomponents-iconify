package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 22V8a4 4 0 0 1 4-4h.5a3.5 3.5 0 0 1 0 7H12h.5A4.5 4.5 0 1 1 8 15.5V15"/>`),
		g.Group(children),
	)
}