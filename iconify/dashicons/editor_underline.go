package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorUnderline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 5h-2v5.71c0 1.99-1.12 2.98-2.45 2.98c-1.32 0-2.55-1-2.55-2.96V5H5v5.87c0 1.91 1 4.54 4.48 4.54c3.49 0 4.52-2.58 4.52-4.5V5zm0 13v-2H5v2h9z"/>`),
		g.Group(children),
	)
}