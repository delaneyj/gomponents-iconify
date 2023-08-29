package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M13.5 10a3.5 3.5 0 0 0-3.5 3.5v73a3.5 3.5 0 0 0 3.5 3.5h73a3.5 3.5 0 0 0 3.5-3.5v-73a3.5 3.5 0 0 0-3.5-3.5h-73zm3.5 7h66v66H17V17z" color="currentColor"/>`),
		g.Group(children),
	)
}