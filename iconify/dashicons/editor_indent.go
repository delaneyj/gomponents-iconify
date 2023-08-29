package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorIndent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 5V3h9v2H3zm10-1V3h4v1h-4zm0 3h2V5l4 3.5l-4 3.5v-2h-2V7zM3 8V6h9v2H3zm2 3V9h7v2H5zm-2 3v-2h9v2H3zm10 0v-1h4v1h-4zm-4 3v-2h3v2H9z"/>`),
		g.Group(children),
	)
}