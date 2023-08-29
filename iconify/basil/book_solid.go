package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 8a4.5 4.5 0 0 1 4.5-4.5h10A1.5 1.5 0 0 1 20 5v15a1 1 0 0 1-1 1H7.5a3.5 3.5 0 0 1-3.465-3H4V8Zm14.5 7.5h-11a2 2 0 1 0 0 4h11v-4ZM8.25 8A.75.75 0 0 1 9 7.25h7a.75.75 0 0 1 0 1.5H9A.75.75 0 0 1 8.25 8ZM9 10.25a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}