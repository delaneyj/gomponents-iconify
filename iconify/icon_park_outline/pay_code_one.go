package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PayCodeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 6h10v10H32zm0 26h10v10H32zM6 32h10v10H6zM6 6h10v10H6zm2 18h22m8 0h2M24 37v2m0-22v14m0-23v2"/>`),
		g.Group(children),
	)
}