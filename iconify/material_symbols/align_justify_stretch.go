package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustifyStretch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22V2h2v20h-2ZM2 22V2h2v20H2Zm11-12V7h5v3h-5Zm-7 0V7h5v3H6Zm7 7v-3h5v3h-5Zm-7 0v-3h5v3H6Z"/>`),
		g.Group(children),
	)
}