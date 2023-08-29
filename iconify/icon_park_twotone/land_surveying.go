package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandSurveying(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLandSurveying0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 26v18m0-18l12 18M24 26L12 44m12-30v6m-5 0h10"/><path fill="#555" d="M10 6h26v8H10z"/><path stroke-linecap="round" d="M40 8v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLandSurveying0)"/>`),
		g.Group(children),
	)
}