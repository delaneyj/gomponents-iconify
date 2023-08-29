package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.031 7.125V9h-2.062l-.004-4H8.016l.027 6H5.984L5.98 7.031H2.035L2.031 9H1.004v1h1.949l.004-2h2.07l.004 4h3.922l.004-6h2.059l.004 4h3.011v1.958L17 9.5l-2.969-2.375z"/>`),
		g.Group(children),
	)
}