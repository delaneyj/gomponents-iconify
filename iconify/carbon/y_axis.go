package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YAxis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 26V7.828l2.586 2.586L12 9L7 4L2 9l1.414 1.414L6 7.828V26a2.002 2.002 0 0 0 2 2h20v-2Z"/>`),
		g.Group(children),
	)
}