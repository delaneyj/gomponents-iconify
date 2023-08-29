package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLetterSpacingWider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h2v16H2Zm18 0V4h2v16h-2ZM7.35 17L11.1 7h1.8l3.75 10h-1.725l-.9-2.55h-4.05l-.9 2.55H7.35Zm3.15-4h3l-1.45-4.15h-.1L10.5 13Z"/>`),
		g.Group(children),
	)
}