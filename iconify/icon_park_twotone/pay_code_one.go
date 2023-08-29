package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PayCodeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPayCodeOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M32 6h10v10H32zm0 26h10v10H32zM6 32h10v10H6zM6 6h10v10H6z"/><path d="M8 24h22m8 0h2M24 37v2m0-22v14m0-23v2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPayCodeOne0)"/>`),
		g.Group(children),
	)
}