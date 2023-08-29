package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongUpE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12.032 1.013l4.21 4.275l-1.424 1.403l-1.804-1.83l-.07 12.116l1.998.01l-.028 6l-6-.029l.029-6l2 .01l.071-12.145L9.161 6.65L7.758 5.224l4.274-4.21Zm-1.108 19.955l2 .01l.01-2l-2-.01l-.01 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}