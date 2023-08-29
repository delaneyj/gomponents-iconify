package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VisualStudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.133 14.242L4.2 18.087l-2.152-1.072V6.994l2.144-1.08l4.9 3.853L16.849 2l5.104 2.033v15.9L16.875 22Zm7.563 1.352V8.406l-4.645 3.598l4.645 3.59ZM4.369 14.301l2.442-2.22l-2.442-2.433Z"/>`),
		g.Group(children),
	)
}