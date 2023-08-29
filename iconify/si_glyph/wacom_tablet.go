package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WacomTablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 5v9h16V5H0zm5 2v5h9V7H5zM.959 7H2v1H.959V7zM2 12H.959v-1H2v1zM.953 10V8.979h2V10h-2zM15 13H4V6h11v7z"/>`),
		g.Group(children),
	)
}