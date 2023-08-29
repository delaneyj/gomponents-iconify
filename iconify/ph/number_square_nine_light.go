package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareNineLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 34H48a14 14 0 0 0-14 14v160a14 14 0 0 0 14 14h160a14 14 0 0 0 14-14V48a14 14 0 0 0-14-14Zm2 174a2 2 0 0 1-2 2H48a2 2 0 0 1-2-2V48a2 2 0 0 1 2-2h160a2 2 0 0 1 2 2ZM145 78.55A34 34 0 1 0 127.94 142a33.56 33.56 0 0 0 5.67-.49l-18.76 31.42a6 6 0 0 0 10.3 6.16l32.3-54.09A34 34 0 0 0 145 78.55Zm2.06 40.45A22 22 0 1 1 139 89a22 22 0 0 1 8.05 30Z"/>`),
		g.Group(children),
	)
}