package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadSideCough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="6" cy="16.999" r="1" fill="currentColor"/><circle cx="2" cy="18" r="1" fill="currentColor"/><circle cx="5" cy="21" r="1" fill="currentColor"/><path fill="currentColor" d="M21.13 21h-8.463a1 1 0 0 1-1-1v-2H10.8a1.935 1.935 0 0 1-1.934-1.934v-1.8H8a1 1 0 0 1-.904-1.426l1.77-3.758v-.016a7.067 7.067 0 0 1 7.284-7.063A7.252 7.252 0 0 1 23 9.321v.212a1.031 1.031 0 0 1-.033.257l-1.796 6.767l.919 3.164A1 1 0 0 1 21.13 21Z" opacity=".5"/>`),
		g.Group(children),
	)
}