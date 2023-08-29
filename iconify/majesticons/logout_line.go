package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoutLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8m4-9l-4-4m4 4l-4 4m4-4H9"/>`),
		g.Group(children),
	)
}