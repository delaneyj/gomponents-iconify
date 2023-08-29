package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplayTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M239 284q0 20-8 30t-23 10t-23-10t-8-29v-17q0-19 8-29t23-10t23 10t8 28v17zm-18-18q0-12-3-17t-10-5t-10 5t-3 15v23q0 11 3 16.5t10 5.5t10-5t3-16v-22zm-74 57h-19v-71l-22 7v-15l39-14h2v93zm24-235q70 0 120 50t50 120.5T291 379t-120.5 50T50 379T0 259h43q0 52 37.5 90t90 38t90.5-38t38-90.5t-38-90t-90-37.5v85L64 109L171 3v85z"/>`),
		g.Group(children),
	)
}