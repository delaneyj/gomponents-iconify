package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.546 15.951c-.509.509-.515-4.982-.515-4.982s-5.493-.007-4.983-.516L10.061.439c.508-.508 1.357-.484 1.896.053l3.55 3.551c.539.539.561 1.388.054 1.896L5.546 15.951z"/>`),
		g.Group(children),
	)
}