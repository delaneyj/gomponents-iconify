package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrailSignSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m502.63 336l-80-80H278v-32h170V64H278V32h-44v32H89.37l-80 80l80 80H234v32H64v160h170v64h44v-64h144.63Z"/>`),
		g.Group(children),
	)
}