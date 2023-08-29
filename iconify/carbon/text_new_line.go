package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextNewLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.586 14.414L24.172 18H6V8H4v10a2.002 2.002 0 0 0 2 2h18.172l-3.586 3.586L22 25l6-6l-6-6Z"/>`),
		g.Group(children),
	)
}