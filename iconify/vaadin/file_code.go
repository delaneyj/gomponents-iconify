package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M6.2 13h-.7l-2-2.5l2-2.5h.7l-2 2.5zm3.6 0h.7l2-2.5l-2-2.5h-.7l2 2.5zm-3.1 1h.6l2.1-7h-.8z"/>`),
		g.Group(children),
	)
}