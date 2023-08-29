package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapSimpleMarkerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 9V4l-6 4l-6-4l-6 4v12l6-4l1 .667"/><path d="M21 15c0 2.611-4 6-4 6s-4-3.389-4-6s1.79-4 4-4s4 1.389 4 4zm-4 .001V15"/></g>`),
		g.Group(children),
	)
}