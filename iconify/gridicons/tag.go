package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2.007h-7.087a2 2 0 0 0-1.414.586l-8.906 8.906a2 2 0 0 0 0 2.828l7.086 7.086a2 2 0 0 0 2.828 0l8.906-8.906c.376-.374.587-.883.587-1.413V4.007a2 2 0 0 0-2-2zM17.007 9a2 2 0 1 1-.001-3.999A2 2 0 0 1 17.007 9z"/>`),
		g.Group(children),
	)
}