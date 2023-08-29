package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RealEstateAgency(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M14.237 39.5H44.72V13.419H14.237V39.5zm15.489-23.485l10.99 9.598h-2.769v11.516h-6.436V29h-4.065v8.129H21.35V25.613h-2.84l11.216-9.598zM10.85 6.984V1.018H4.076V50h6.774V10.033h35.226V6.984z"/>`),
		g.Group(children),
	)
}