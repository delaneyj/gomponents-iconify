package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.934 15.966L22.9 16L12.267 5.392L1.633 16l-.034-.034v-5.319L12.266 0l10.666 10.647z"/><path fill="currentColor" d="M10.49 24V5.334h3.555V24z"/>`),
		g.Group(children),
	)
}