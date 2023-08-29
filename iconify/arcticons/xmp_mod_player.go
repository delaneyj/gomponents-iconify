package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XmpModPlayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 10.419H28.269L24 15.603l-4.269-5.184H5.5l11.385 13.824L5.5 38.067h14.231L24 32.883l10.063 12.198l8.437-7.014l-11.385-13.824L42.5 10.419z"/>`),
		g.Group(children),
	)
}