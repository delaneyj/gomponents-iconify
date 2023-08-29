package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtitlesOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 10h-5.15l-6-6H22v15.15L14.85 12H18v-2Zm2.55 13.35L17.15 20H2V4h2v2.8L.65 3.45l1.4-1.4l19.9 19.9l-1.4 1.4ZM6 12h2v-1.2l-.8-.8H6v2Zm5.15 2H6v2h7.15l-2-2Z"/>`),
		g.Group(children),
	)
}