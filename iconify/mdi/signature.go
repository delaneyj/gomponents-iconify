package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 22H2v-2h20v2M2.26 16.83L5.09 14l-2.83-2.83l1.41-1.41l2.83 2.83l2.83-2.83l1.41 1.41L7.91 14l2.83 2.83l-1.41 1.41l-2.83-2.83l-2.83 2.83l-1.41-1.41Z"/>`),
		g.Group(children),
	)
}