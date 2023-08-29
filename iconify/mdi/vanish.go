package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vanish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13v-2h5v2h-5m-1.17-5.24l2.83-2.83l1.41 1.41l-2.83 2.83l-1.41-1.41M11 16h2v5h-2v-5m0-13h2v5h-2V3M4.93 17.66l2.83-2.83l1.41 1.41l-2.83 2.83l-1.41-1.41m0-11.32l1.41-1.41l2.83 2.83l-1.41 1.41l-2.83-2.83M8 13H3v-2h5v2m11.07 4.66l-1.41 1.41l-2.83-2.83l1.41-1.41l2.83 2.83Z"/>`),
		g.Group(children),
	)
}