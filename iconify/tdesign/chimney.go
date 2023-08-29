package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chimney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.719 2H13.28l-1.27 5.08l.358 4.118L15 9.882l7 3.5V22H1.927L2.99 7.088L1.719 2ZM4.93 8l-.857 12H8v-6.618l2.444-1.223L10.083 8H4.93Zm5.288-2l.5-2H4.28l.5 2h5.439Zm-.22 14h4v-3h2v3h4v-5.382l-5-2.5l-5 2.5V20Z"/>`),
		g.Group(children),
	)
}