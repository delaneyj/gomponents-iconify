package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresentationReport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-3.586l2.293 2.293a1 1 0 0 1-1.414 1.414L12 17.414l-3.293 3.293a1 1 0 0 1-1.414-1.414L9.586 17H6a3 3 0 0 1-3-3V6zm14 2a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0V8zm-4 2a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2zm-4 1a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0v-1z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}