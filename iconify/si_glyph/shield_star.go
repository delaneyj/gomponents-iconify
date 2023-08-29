package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 0C4.688 0 2 1.958 2 1.958l.005 6.098C2.005 13.091 8.002 16 8.002 16S14 13.259 14 8.041V1.958S11.314 0 8 0zm1.607 7.875l.955 2.939l-2.5-1.816l-2.502 1.816l.955-2.939l-2.5-1.816h3.091l.956-2.94l.955 2.94h3.091L9.607 7.875z"/>`),
		g.Group(children),
	)
}