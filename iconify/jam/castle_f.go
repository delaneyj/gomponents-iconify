package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 8v11h-3v-4a2 2 0 1 0-4 0v4H0v-8a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1h2V8a2 2 0 0 1-2-2V1a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1h2V1a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v5a2 2 0 0 1-2 2zm-6 1a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2z"/>`),
		g.Group(children),
	)
}