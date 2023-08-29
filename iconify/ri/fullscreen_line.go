package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3v2H4v4H2V3h6ZM2 21v-6h2v4h4v2H2Zm20 0h-6v-2h4v-4h2v6Zm0-12h-2V5h-4V3h6v6Z"/>`),
		g.Group(children),
	)
}