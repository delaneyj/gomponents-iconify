package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePresentation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zm3 15H3V1h6v4h4v10zM10 4V1l3 3h-3z"/><path fill="currentColor" d="M9 6H7v1H4v6h2v1h1v-1h2v1h1v-1h2V7H9V6zm2 2v4H5V8h6z"/><path fill="currentColor" d="M7 9v2l2-1z"/>`),
		g.Group(children),
	)
}