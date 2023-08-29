package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessPawn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.646 9a4 4 0 1 0-5.292 0H13a2 2 0 1 0 0 4h.5v4s0 .5-.5 1l-2.025 2.025c-1.493 1.493-1.26 3.703.015 4.975A3 3 0 0 0 8 28v2h16v-2a3 3 0 0 0-2.99-3c1.274-1.272 1.508-3.482.015-4.975L19 18c-.5-.5-.5-1-.5-1v-4h.5a2 2 0 1 0 0-4h-.354Z"/>`),
		g.Group(children),
	)
}