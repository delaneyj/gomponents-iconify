package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RawOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15V9h3.5q.6 0 1.05.45T8 10.5v1q0 .45-.238.813T7.1 12.9L8 15H6.5l-.9-2H4.5v2H3Zm5.75 0l1.5-6h2.5l1.5 6h-1.5l-.35-1.5h-1.75l-.4 1.5h-1.5Zm6.75 0L14 9h1.5l.75 3L17 9h1.5l.75 3L20 9h1.5L20 15h-1.5l-.75-3.05L17 15h-1.5ZM11 12h1l-.25-1h-.5L11 12Zm-6.5-.5h2v-1h-2v1Z"/>`),
		g.Group(children),
	)
}