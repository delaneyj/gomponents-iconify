package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Captions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11 8v2H8v4h3v2H6V8h5Zm7 0v2h-3v4h3v2h-5V8h5Z"/><path fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V5Zm2 13V6h16v12H4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}