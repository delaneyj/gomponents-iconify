package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavaidVor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="16" r="2" fill="currentColor"/><path fill="currentColor" d="m30.864 15.496l-7-12A1 1 0 0 0 23 3H9a1 1 0 0 0-.864.496l-7 12a1 1 0 0 0 0 1.008l7 12A1 1 0 0 0 9 29h14a1 1 0 0 0 .864-.496l7-12a1 1 0 0 0 0-1.008ZM22.426 27H9.574L3.158 16L9.574 5h12.852l6.416 11Z"/>`),
		g.Group(children),
	)
}