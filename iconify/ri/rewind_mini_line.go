package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindMiniLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9.86L5.968 12L9 14.14V9.86Zm1.909 7.463a.5.5 0 0 1-.697.12l-7.133-5.035a.5.5 0 0 1 0-.817l7.133-5.035a.5.5 0 0 1 .788.409v10.07a.5.5 0 0 1-.091.288ZM18 14.14V9.86L14.968 12L18 14.14Zm-5.921-1.732a.5.5 0 0 1 0-.817l7.133-5.035a.5.5 0 0 1 .788.409v10.07a.5.5 0 0 1-.788.408l-7.133-5.035Z"/>`),
		g.Group(children),
	)
}