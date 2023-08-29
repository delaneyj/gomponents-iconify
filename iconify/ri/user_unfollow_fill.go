package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserUnfollowFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 14.252V22H4a8 8 0 0 1 10-7.748ZM12 13c-3.315 0-6-2.685-6-6s2.685-6 6-6s6 2.685 6 6s-2.685 6-6 6Zm7 3.586l2.121-2.121l1.415 1.414L20.413 18l2.121 2.121l-1.414 1.415L19 19.413l-2.121 2.121l-1.415-1.414L17.587 18l-2.121-2.121l1.414-1.415L19 16.587Z"/>`),
		g.Group(children),
	)
}