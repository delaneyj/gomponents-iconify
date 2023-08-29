package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 11c-.2 0-.4 0-.5.1L17.4 9H20V6l-3.7 1.9L13.4 5H9v2h3.6l2 2H11l-4 2l-2-2H0v2h4c-2.2 0-4 1.8-4 4s1.8 4 4 4s4-1.8 4-4l2 2h3l3.5-6.1l1 1c-.9.7-1.5 1.9-1.5 3.1c0 2.2 1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4M4 17c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2m16 0c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2Z"/>`),
		g.Group(children),
	)
}