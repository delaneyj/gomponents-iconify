package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AxisZRotateCounterclockwise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 12l4 4l-4 4v-3.1c-4.56-.46-8-2.48-8-4.9c0-2.42 3.44-4.44 8-4.9v1.99C6.55 9.43 4 10.6 4 12c0 1.4 2.55 2.57 6 2.91V12m10 0c0-1.4-2.55-2.57-6-2.91V7.1c4.56.46 8 2.48 8 4.9c0 2.16-2.74 4-6.58 4.7l.7-.7l-1.2-1.21C17.89 14.36 20 13.27 20 12M11 2h2v11l-2-2V2m0 20v-1l2-2v3h-2Z"/>`),
		g.Group(children),
	)
}