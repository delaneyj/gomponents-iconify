package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.984 0L2.016.052l.01 9.927l4.005 5.037H4v.953h1.031V16h8.938V6.012H7.984V0z"/><path d="M9 0v4.969h5L9 0zM1.016 13.953h.993v2.006h.966v-2.006h.978l-1.422-1.937l-1.515 1.937z"/></g>`),
		g.Group(children),
	)
}