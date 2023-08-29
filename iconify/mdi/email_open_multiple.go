package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailOpenMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 8l8 5l8-5l-8-5l-8 5m18 0v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V8c0-.73.39-1.36.97-1.71L14 .64l9.03 5.65c.58.35.97.98.97 1.71M2 8v14h18v2H2a2 2 0 0 1-2-2V8h2Z"/>`),
		g.Group(children),
	)
}