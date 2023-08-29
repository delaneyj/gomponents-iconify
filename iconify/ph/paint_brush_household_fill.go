package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintBrushHouseholdFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230.64 25.36a32 32 0 0 0-45.26 0c-.1.1-.2.2-.29.31l-54.91 61.18l-9.18-9.21a24 24 0 0 0-33.95 0l-76.69 76.7a8 8 0 0 0 0 11.31l80 80a8 8 0 0 0 11.31 0L178.36 169a24 24 0 0 0 0-33.95l-9.21-9.2l61.18-54.91a2.91 2.91 0 0 0 .31-.3a32 32 0 0 0 0-45.28ZM96 228.69L79.32 212l22.34-22.35a8 8 0 0 0-11.31-11.31L68 200.68L55.32 188l22.34-22.35a8 8 0 0 0-11.31-11.31L44 176.68L27.31 160L72 115.31L140.69 184Z"/>`),
		g.Group(children),
	)
}