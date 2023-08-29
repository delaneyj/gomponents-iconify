package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Factory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 2v7L9 6v3L5 6v3L0 6v9h15V2h-2zM3 12H1v-2h2v2zm4 0H5v-2h2v2zm4 0H9v-2h2v2z"/>`),
		g.Group(children),
	)
}