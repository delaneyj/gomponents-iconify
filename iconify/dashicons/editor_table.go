package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 17V3H2v14h16zM16 7H4V5h12v2zm-7 4H4V9h5v2zm7 0h-5V9h5v2zm-7 4H4v-2h5v2zm7 0h-5v-2h5v2z"/>`),
		g.Group(children),
	)
}