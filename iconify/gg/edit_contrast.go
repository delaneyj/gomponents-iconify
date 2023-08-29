package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditContrast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 18a5.978 5.978 0 0 1-4-1.528A5.985 5.985 0 0 1 6 12c0-1.777.772-3.374 2-4.472A5.978 5.978 0 0 1 12 6v12Z"/><path fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm10 8a8 8 0 1 1 0-16a8 8 0 0 1 0 16Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}