package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrthogonalView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 21V3h18v18H3Zm0-4.5h18M3 12h18M3 7.5h18M16.5 3v18M12 3v18M7.5 3v18"/>`),
		g.Group(children),
	)
}