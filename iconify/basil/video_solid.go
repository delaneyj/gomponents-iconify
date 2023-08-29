package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.966 5.814c2.353-.21 4.72-.192 7.07.055l.608.064A2.595 2.595 0 0 1 16.93 8.08l3.279-1.742a.75.75 0 0 1 1.099.597l.025.283a54.758 54.758 0 0 1 0 9.564l-.025.284a.75.75 0 0 1-1.1.596l-3.278-1.741a2.596 2.596 0 0 1-2.287 2.146l-.608.064c-2.35.247-4.717.265-7.07.055l-1.581-.141a2.585 2.585 0 0 1-2.328-2.201a26.568 26.568 0 0 1 0-7.687a2.585 2.585 0 0 1 2.328-2.201l1.581-.142Zm10.15 8.553a.75.75 0 0 1 .236.078l2.555 1.358a53.246 53.246 0 0 0 0-7.606l-2.555 1.358a.749.749 0 0 1-.236.079a26.72 26.72 0 0 1 0 4.733Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}