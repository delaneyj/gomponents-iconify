package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneGetApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.17 11H13V5h-2v6H9.83L12 13.17z" opacity=".3"/><path fill="currentColor" d="M19 9h-4V3H9v6H5l7 7l7-7zm-8 2V5h2v6h1.17L12 13.17L9.83 11H11zm-6 7h14v2H5z"/>`),
		g.Group(children),
	)
}