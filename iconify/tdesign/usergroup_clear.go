package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsergroupClear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 5a3 3 0 0 0 0 6v2a5 5 0 0 0-5 5v4H2v-4a7 7 0 0 1 3.75-6.2A5 5 0 0 1 9 3h1v2H9Zm6 0a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0ZM7 19a5 5 0 0 1 5-5h3v2h-3a3 3 0 0 0-3 3v1h6v2H7v-3Zm10.879-4.536L20 16.586l2.121-2.122l1.415 1.415L21.413 18l2.121 2.121l-1.414 1.415L20 19.413l-2.121 2.121l-1.415-1.414L18.587 18l-2.121-2.121l1.414-1.415Z"/>`),
		g.Group(children),
	)
}