package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLetterSpacingWide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20V4h2v16H3Zm16 0V4h2v16h-2ZM7.35 17L11.1 7h1.8l3.75 10H14.9l-.85-2.55H10L9.1 17H7.35Zm3.15-4h3l-1.45-4.15L10.5 13Z"/>`),
		g.Group(children),
	)
}