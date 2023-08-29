package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutSidebarRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 5H2v14h20V5zm-2 2v10h-2V7h2zm-4 0v10H4V7h12z"/>`),
		g.Group(children),
	)
}