package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1V3Zm2 2v14h18V5H3Zm1.996 2.5H7v2.004H4.996V7.5Zm4 0H11v2.004H8.996V7.5Zm4 0H15v2.004h-2.004V7.5Zm4 0H19v2.004h-2.004V7.5Zm-12 3H7v2.004H4.996V10.5Zm4 0H11v2.004H8.996V10.5Zm4 0H15v2.004h-2.004V10.5Zm4 0H19v2.004h-2.004V10.5ZM5 15h14v2H5v-2Z"/>`),
		g.Group(children),
	)
}