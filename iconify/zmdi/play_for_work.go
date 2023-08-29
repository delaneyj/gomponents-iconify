package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayForWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M107 43h42v119h75l-96 96l-96-96h75V43zM0 235h43q0 35 25 60t60 25t60-25t25-60h43q0 53-37.5 90.5T128 363t-90.5-37.5T0 235z"/>`),
		g.Group(children),
	)
}