package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserErrorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 4a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0Zm13.879-.536L19.5 8.586l2.121-2.122l1.415 1.415L20.913 10l2.121 2.121l-1.414 1.415l-2.121-2.122l-2.121 2.122l-1.415-1.415L18.087 10l-2.121-2.121l1.414-1.415ZM0 19a5 5 0 0 1 5-5h7a5 5 0 0 1 5 5v2h-2v-2a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v2H0v-2Z"/>`),
		g.Group(children),
	)
}