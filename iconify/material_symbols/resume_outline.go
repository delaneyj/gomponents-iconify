package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResumeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18V6h2v12H6Zm4 0l10-6l-10-6v12Zm2-3.525v-4.95L16.125 12L12 14.475ZM12 12Z"/>`),
		g.Group(children),
	)
}