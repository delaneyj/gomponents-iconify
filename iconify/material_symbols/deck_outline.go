package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22V9H2l10-7l10 7h-9v13h-2Zm1-15h3.65h-7.3H12ZM3 22v-5.25l-.8-4.4L4.15 12l.75 4H9v6H7v-4H5v4H3Zm12 0v-6h4.1l.75-4l1.95.35l-.8 4.4V22h-2v-4h-2v4h-2ZM8.35 7h7.3L12 4.45L8.35 7Z"/>`),
		g.Group(children),
	)
}