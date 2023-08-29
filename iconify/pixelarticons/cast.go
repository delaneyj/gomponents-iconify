package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3h18v18h-8v-2h6V5H4v4H2V3h2zm0 16H2v2h2v-2zm-2-4h4v2H2v-2zm8-4H2v2h8v8h2V11h-2zm-4 4h2v6H6v-6z"/>`),
		g.Group(children),
	)
}