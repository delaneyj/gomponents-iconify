package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m13.086 7l-2.39-2.398l.702-.704L15 7.5l-3.602 3.602l-.703-.704l2.383-2.382V8H3V7h10.086zM1 4h1v7H1V4z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}