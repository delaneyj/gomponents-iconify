package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m9 13.887l5-5V8.18l-5-5l-.707.707l4.146 4.147H2v1h10.44L8.292 13.18l.707.707z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}