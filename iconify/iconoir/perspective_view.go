package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerspectiveView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1 21L4.143 3h15.714L23 21H1Zm1-4.5h20M3 12h18M4 7.5h16M12 3v18M8 3.5l-1.5 17m9.5-17l1.5 17"/>`),
		g.Group(children),
	)
}