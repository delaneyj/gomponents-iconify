package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3h341zM128 387v-86H43v86h85zm0-128v-86H43v86h85zm0-128V45H43v86h85zm128 256v-86h-85v86h85zm0-128v-86h-85v86h85zm0-128V45h-85v86h85zm128 256v-86h-85v86h85zm0-128v-86h-85v86h85zm0-128V45h-85v86h85z"/>`),
		g.Group(children),
	)
}