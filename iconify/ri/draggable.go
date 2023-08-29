package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Draggable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5 5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM15.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5 5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm-1.5 8a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}