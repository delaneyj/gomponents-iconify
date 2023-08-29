package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunctionMath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 18h-2l-2 3.897L20 18h-2l2.905 5L18 28h2l2-3.799L24 28h2l-2.902-5L26 18zM19 6V4h-5.087a1.99 1.99 0 0 0-1.992 1.819L11.27 13H7v2h4.087l-1 11H5v2h5.087a1.99 1.99 0 0 0 1.992-1.819L13.095 15H18v-2h-4.723l.636-7z"/>`),
		g.Group(children),
	)
}