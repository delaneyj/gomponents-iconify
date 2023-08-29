package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.977 2.806a1.692 1.692 0 0 0-2.304.082L3.28 6.28a.75.75 0 1 1-1.06-1.06l3.392-3.392a3.192 3.192 0 1 1 4.457 4.569l-4.723 4.496A1.923 1.923 0 0 1 2.66 8.14l4.243-4.244a.75.75 0 1 1 1.06 1.061l-4.24 4.244a.423.423 0 0 0 .59.605L9.035 5.31a1.692 1.692 0 0 0-.058-2.504Z"/>`),
		g.Group(children),
	)
}