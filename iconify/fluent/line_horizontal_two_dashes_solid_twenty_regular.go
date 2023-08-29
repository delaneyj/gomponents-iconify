package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineHorizontalTwoDashesSolidTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 6a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3Zm6 0a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3Zm6 0a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3Zm-12 6a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1h-15Z"/>`),
		g.Group(children),
	)
}