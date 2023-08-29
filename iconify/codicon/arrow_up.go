package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m13.854 7l-5-5h-.707l-5 5l.707.707L8 3.561V14h1V3.56l4.146 4.147l.708-.707z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}