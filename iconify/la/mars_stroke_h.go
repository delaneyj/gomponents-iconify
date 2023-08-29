package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarsStrokeH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 8c-4.406 0-8 3.594-8 8c0 4.406 3.594 8 8 8c4.066 0 7.438-3.066 7.938-7H20v3h2v-3h4.563l-4.282 4.281l1.438 1.438l6-6l.687-.719l-.687-.719l-6-6l-1.438 1.438L26.563 15H22v-3h-2v3h-2.063c-.5-3.934-3.87-7-7.937-7zm0 2c3.324 0 6 2.676 6 6s-2.676 6-6 6s-6-2.676-6-6s2.676-6 6-6z"/>`),
		g.Group(children),
	)
}