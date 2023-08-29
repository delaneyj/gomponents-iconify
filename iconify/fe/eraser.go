package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eraser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feEraser0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feEraser1" fill="currentColor"><path id="feEraser2" d="M18 4v16a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2ZM8 4v4h8V4H8Z"/></g></g>`),
		g.Group(children),
	)
}