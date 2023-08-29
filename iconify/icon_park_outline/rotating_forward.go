package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotatingForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 6.676C8.022 10.134 4 16.597 4 24M14 6.676V14m0-7.324H6.676m0 27.324C10.134 39.978 16.597 44 24 44M6.676 34H14m-7.324 0v7.324m27.324 0C39.978 37.866 44 31.403 44 24M34 41.324V34m0 7.324h7.324m0-27.324C37.866 8.022 31.403 4 24 4m17.324 10H34m7.324 0V6.676"/>`),
		g.Group(children),
	)
}