package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDownTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.794 7.733a.75.75 0 0 1-.026 1.06l-5.25 5.001a.75.75 0 0 1-1.035 0l-5.25-5a.75.75 0 0 1 1.034-1.087l4.734 4.509l4.733-4.51a.75.75 0 0 1 1.06.027Z"/>`),
		g.Group(children),
	)
}