package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23.707 18.707a1 1 0 0 1-1.414 0L14 10.414l-8.293 8.293a1 1 0 0 1-1.414-1.414l9-9a1 1 0 0 1 1.414 0l9 9a1 1 0 0 1 0 1.414Z"/>`),
		g.Group(children),
	)
}