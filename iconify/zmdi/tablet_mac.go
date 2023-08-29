package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletMac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M352 0q22 0 37.5 15.5T405 53v406q0 22-15.5 37.5T352 512H53q-22 0-37.5-15.5T0 459V53q0-22 15.5-37.5T53 0h299zM202.5 491q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5zM363 405V64H43v341h320z"/>`),
		g.Group(children),
	)
}