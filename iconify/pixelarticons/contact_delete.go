package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 3H0v18h16v-2H2V5h20v10h2V3h-2zM6 7h4v4H6V7zm0 8H4v2h2v-2zm4 0H6v-2h4v2zm0 0v2h2v-2h-2zm4-8h6v2h-6V7zm6 4h-6v2h6v-2zm-6 4h2v2h-2v-2zm8 4h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2zm0 0h2v-2h-2v2z"/>`),
		g.Group(children),
	)
}