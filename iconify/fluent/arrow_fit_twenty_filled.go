package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFitTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.791 6.72a.75.75 0 0 1-.002 1.061L4.566 9H8.25a.75.75 0 0 1 0 1.5H4.56l1.22 1.22a.75.75 0 1 1-1.06 1.06l-2.5-2.5a.75.75 0 0 1 0-1.061l2.51-2.5a.75.75 0 0 1 1.061.002Zm8.429 0a.75.75 0 0 1 1.06 0l2.5 2.5a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 1 1-1.06-1.06l1.22-1.22h-3.69a.75.75 0 0 1 0-1.5h3.69l-1.22-1.22a.75.75 0 0 1 0-1.06Z"/>`),
		g.Group(children),
	)
}