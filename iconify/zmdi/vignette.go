package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vignette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h384zM234.5 320q70.5 0 120.5-37.5t50-90.5t-50-90.5T234.5 64T114 101.5T64 192t50 90.5T234.5 320z"/>`),
		g.Group(children),
	)
}