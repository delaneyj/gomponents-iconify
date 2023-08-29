package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtitlesOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h2l2 2H4v12h11.15l-2-2H6v-2h5.15L.65 3.45l1.4-1.4l19.9 19.9l-1.4 1.4l-3.4-3.35H2Zm18-2.85V6H8.85l-2-2H22v15.15l-2-2ZM14.85 12l-2-2H18v2h-3.15ZM6 12v-2h2v2H6Zm8.425-.425Zm-4.85.85Z"/>`),
		g.Group(children),
	)
}