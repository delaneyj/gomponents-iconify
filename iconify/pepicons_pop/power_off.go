package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M4.338 5.179a1 1 0 1 1 1.424 1.404A6.127 6.127 0 0 0 4 10.901C4 14.272 6.69 17 10 17s6-2.728 6-6.1c0-1.643-.641-3.18-1.762-4.317a1 1 0 1 1 1.424-1.404A8.127 8.127 0 0 1 18 10.9c0 4.47-3.579 8.1-8 8.1c-4.421 0-8-3.63-8-8.1c0-2.173.85-4.213 2.338-5.721Z"/><rect width="2" height="10.5" x="9" y="1" rx="1"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}