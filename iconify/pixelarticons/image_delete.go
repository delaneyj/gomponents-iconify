package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 3H2v18h20V11h-2v8H4V5h10V3zM6 7h2v2H6V7zm14-2h-2V3h-2v2h2v2h-2v2h2V7h2v2h2V7h-2V5zm0 0V3h2v2h-2zm-8 4h2v2h-2V9zm-2 4v-2h2v2h-2zm-2 2h2v-2H8v2zm0 0v2H6v-2h2zm8-2h-2v-2h2v2zm0 0h2v2h-2v-2z"/>`),
		g.Group(children),
	)
}