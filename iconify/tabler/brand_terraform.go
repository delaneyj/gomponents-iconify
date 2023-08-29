package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandTerraform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 15.5L3.524 9.284A1 1 0 0 1 3 8.404V4.35a1.35 1.35 0 0 1 2.03-1.166L15 9v10.65a1.35 1.35 0 0 1-2.03 1.166l-3.474-2.027A1 1 0 0 1 9 17.926V6m6 9.5l5.504-3.21a1 1 0 0 0 .496-.864V7.85a1.35 1.35 0 0 0-2.03-1.166L15 9"/>`),
		g.Group(children),
	)
}