package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 6h-4V4h6V2H9v2h6v2h-4a5.006 5.006 0 0 0-5 5v11a4.99 4.99 0 0 0 3.582 4.77L8.198 30h2.176l1.285-3h8.682l1.286 3h2.175l-1.384-3.23A4.99 4.99 0 0 0 26 22V11a5.006 5.006 0 0 0-5-5ZM11 8h10a2.995 2.995 0 0 1 2.816 2H8.184A2.995 2.995 0 0 1 11 8Zm13 13h-3v2h2.816A2.995 2.995 0 0 1 21 25H11a2.995 2.995 0 0 1-2.816-2H11v-2H8v-2h16Zm0-4H8v-5h16Z"/>`),
		g.Group(children),
	)
}