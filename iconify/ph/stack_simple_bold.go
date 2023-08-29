package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m10.05 110.42l112 64a12 12 0 0 0 11.9 0l112-64a12 12 0 0 0 0-20.84l-112-64a12 12 0 0 0-11.9 0l-112 64a12 12 0 0 0 0 20.84Zm118-60.6L215.81 100L128 150.18L40.19 100Zm122.42 92.23a12 12 0 0 1-4.47 16.37l-112 64a12 12 0 0 1-11.9 0l-112-64A12 12 0 1 1 22 137.58l106 60.6l106-60.6a12 12 0 0 1 16.42 4.47Z"/>`),
		g.Group(children),
	)
}