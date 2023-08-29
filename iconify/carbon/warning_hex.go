package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningHex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 21a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 0 0 16 21zM15 8h2v10h-2z"/><path fill="currentColor" d="M23 29H9a1 1 0 0 1-.864-.496l-7-12a1 1 0 0 1 0-1.008l7-12A1 1 0 0 1 9 3h14a1 1 0 0 1 .864.496l7 12a1 1 0 0 1 0 1.008l-7 12A1 1 0 0 1 23 29ZM9.574 27h12.852l6.416-11l-6.416-11H9.574L3.158 16Z"/>`),
		g.Group(children),
	)
}