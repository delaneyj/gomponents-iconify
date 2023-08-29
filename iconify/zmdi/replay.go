package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Replay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 91q70 0 120 50t50 120.5T291 382t-120.5 50T50 382T0 261h43q0 53 37.5 90.5T171 389t90.5-37.5T299 261t-37.5-90.5T171 133v86L64 112L171 5v86z"/>`),
		g.Group(children),
	)
}