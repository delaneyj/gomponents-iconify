package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperBoardSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 8.5a.5.5 0 1 0 0-1H13V6h1.5a.5.5 0 1 0 0-1H13a2 2 0 0 0-2-2V1.5a.5.5 0 1 0-1 0V3H8.5V1.5a.5.5 0 1 0-1 0V3H6V1.5a.5.5 0 1 0-1 0V3a2 2 0 0 0-2 2H1.5a.5.5 0 1 0 0 1H3v1.5H1.5a.5.5 0 1 0 0 1H3V10H1.5a.5.5 0 1 0 0 1H3a2 2 0 0 0 2 2v1.5a.5.5 0 1 0 1 0V13h1.5v1.5a.5.5 0 1 0 1 0V13H10v1.5a.5.5 0 1 0 1 0V13a2 2 0 0 0 2-2h1.5a.5.5 0 1 0 0-1H13V8.5h1.5Zm-6.5 2A2.502 2.502 0 0 1 5.5 8c0-1.379 1.121-2.5 2.5-2.5s2.5 1.121 2.5 2.5s-1.121 2.5-2.5 2.5Zm0-4c-.827 0-1.5.673-1.5 1.5S7.173 9.5 8 9.5S9.5 8.827 9.5 8S8.827 6.5 8 6.5Z"/>`),
		g.Group(children),
	)
}