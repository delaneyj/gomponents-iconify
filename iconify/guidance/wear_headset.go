package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WearHeadset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M19.819 16.5c-.375.94-.937 1.828-1.767 2.735c-1.702 1.86-3.94 3.452-6.052 4.015c-2.109-.562-4.342-2.15-6.042-4.006c-.835-.911-1.4-1.801-1.776-2.744M3.5 7.944L1 7.5s-.5 1.436-.5 4s.5 4 .5 4l2.5-.444V7.944Zm0 0V4.5s2.551-4 8.5-4s8.5 4 8.5 4v3.444m0 0L23 7.5s.5 1.436.5 4s-.5 4-.5 4l-2.5-.444V7.944ZM5 5s2 1.5 5.5 1.5L9.5 4c2 2 7.5 2.286 9.5 2.286"/>`),
		g.Group(children),
	)
}