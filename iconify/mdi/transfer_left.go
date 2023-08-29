package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransferLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 16h-2V8h2v8m-4 0h-2V8h2v8m-4 0h-2V8h2v8M9 5v14l-7-7l7-7Z"/>`),
		g.Group(children),
	)
}