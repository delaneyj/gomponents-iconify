package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExpandLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9 4a1 1 0 0 1-1 1H6.414l3.293 3.293a1 1 0 0 1-1.414 1.414L5 6.414V8a1 1 0 1 1-2 0V4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1zm8.586 15l-3.293-3.293a1 1 0 0 1 1.414-1.414L19 17.586V16a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1h-4a1 1 0 1 1 0-2h1.586zM9 20a1 1 0 0 0-1-1H6.414l3.293-3.293a1 1 0 0 0-1.414-1.414L5 17.586V16a1 1 0 1 0-2 0v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1zm8.586-15l-3.293 3.293a1 1 0 0 0 1.414 1.414L19 6.414V8a1 1 0 1 0 2 0V4a1 1 0 0 0-1-1h-4a1 1 0 1 0 0 2h1.586z"/></g>`),
		g.Group(children),
	)
}