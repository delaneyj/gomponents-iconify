package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PointDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 6.5V9a1.502 1.502 0 0 1-2.236 1.307a1.5 1.5 0 0 1-2.264.31a1.494 1.494 0 0 1-1.5.297V14.5c0 .827-.673 1.5-1.5 1.5S6 15.327 6 14.5V8.333L3.25 9.799a1.502 1.502 0 0 1-1.789-2.381l.012-.011L5.21 4H4.5a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-.691l1.138 2.276A.496.496 0 0 1 15 6.5zm-1-4a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0zm0 4.118L12.691 4H6.694L2.15 8.143a.5.5 0 0 0 .614.782l3.5-1.866a.502.502 0 0 1 .735.442v7a.5.5 0 0 0 1 0v-5a.5.5 0 0 1 1 0a.5.5 0 0 0 1 0a.5.5 0 0 1 1 0a.5.5 0 0 0 1 0v-.5a.5.5 0 0 1 1 0a.5.5 0 0 0 1 0V6.619z"/>`),
		g.Group(children),
	)
}