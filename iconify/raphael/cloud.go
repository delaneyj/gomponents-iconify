package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.345 13.904a6.25 6.25 0 0 0-6.22-6.84a6.244 6.244 0 0 0-5.847 4.058a3.205 3.205 0 0 0-2.372-1.06a3.22 3.22 0 0 0-3.22 3.22c0 .21.024.415.063.613c-2.373.39-4.188 2.436-4.188 4.918a5 5 0 0 0 5 5h15.875a4.998 4.998 0 0 0 .908-9.91z"/>`),
		g.Group(children),
	)
}