package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 0q53 0 90.5 37.5T277 128t-37.5 90.5T149 256H85v128H0V0h149zm5 171q17 0 29.5-12.5T196 128t-12.5-30.5T154 85H85v86h69z"/>`),
		g.Group(children),
	)
}