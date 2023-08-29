package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feUnderline0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUnderline1" fill="currentColor"><path id="feUnderline2" d="M17 21H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Zm1-17v7a6 6 0 1 1-12 0V4a1 1 0 1 1 2 0v7c0 2.21 1.79 4 4 4s4-1.79 4-4V4a1 1 0 0 1 2 0Z"/></g></g>`),
		g.Group(children),
	)
}