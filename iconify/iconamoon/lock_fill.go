package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7a3 3 0 1 1 6 0v2H9V7ZM7 9V7a5 5 0 0 1 10 0v2h2a1 1 0 0 1 1 1v9a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3v-9a1 1 0 0 1 1-1h2Zm6 6.5a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5h-.01a1.5 1.5 0 0 1-1.5-1.5v-.01Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}