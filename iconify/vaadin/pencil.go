package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 11.9L0 16l4.1-1l9.2-9.2l-3.1-3.1L1 11.9zm.5 3.1l-.4-.5l.4-2l2 2l-2 .5zm9.4-10.6l-8.1 8l-.6-.6l8.1-8l.6.6zM15.3.7C14.2-.4 12.7.2 12.7.2l-1.5 1.5l3.1 3.1l1.5-1.5c0-.1.6-1.5-.5-2.6zm-1.9.9l-.5-.5C13.5.5 14 1 14 1l-.6.6z"/>`),
		g.Group(children),
	)
}