package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Integration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.88 15.52l-6-11A1 1 0 0 0 23 4H9a1 1 0 0 0-.88.52l-6 11a1 1 0 0 0 0 1l6 11A1 1 0 0 0 9 28h14a1 1 0 0 0 .88-.52l6-11a1 1 0 0 0 0-.96ZM22.93 7l4.39 8h-9.5ZM16 14.14L10.82 6h10.36ZM9.07 7l5.11 8h-9.5ZM4.68 17h9.5l-5.11 8Zm11.32.86L21.18 26H10.82ZM22.93 25l-5.11-8h9.5Z"/>`),
		g.Group(children),
	)
}