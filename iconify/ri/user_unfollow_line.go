package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserUnfollowLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 14.252v2.09A6 6 0 0 0 6 22H4a8 8 0 0 1 10-7.749ZM12 13c-3.315 0-6-2.685-6-6s2.685-6 6-6s6 2.685 6 6s-2.685 6-6 6Zm0-2c2.21 0 4-1.79 4-4s-1.79-4-4-4s-4 1.79-4 4s1.79 4 4 4Zm7 6.586l2.121-2.121l1.415 1.414L20.413 19l2.121 2.121l-1.414 1.415L19 20.413l-2.121 2.121l-1.415-1.414L17.587 19l-2.121-2.121l1.414-1.415L19 17.587Z"/>`),
		g.Group(children),
	)
}