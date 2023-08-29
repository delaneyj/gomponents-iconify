package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zm3 15H3V1h6v4h4v10zM10 4V1l3 3h-3z"/><path fill="currentColor" d="M5 6v6l6-3z"/>`),
		g.Group(children),
	)
}