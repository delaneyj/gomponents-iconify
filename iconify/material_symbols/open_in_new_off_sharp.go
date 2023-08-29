package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInNewOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L18.15 21H3V5.85L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425ZM16.15 19l-4.875-4.875L9.7 15.7l-1.4-1.4l1.575-1.575L5 7.85V19h11.15ZM7.85 5l-2-2H12v2H7.85Zm6.275 6.275l-1.4-1.4L17.6 5H14V3h7v7h-2V6.4l-4.875 4.875ZM21 18.15l-2-2V12h2v6.15Z"/>`),
		g.Group(children),
	)
}