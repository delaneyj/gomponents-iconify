package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistRepeatListFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 5a1 1 0 0 1 1.707-.707l2 2a1 1 0 0 1 0 1.414l-2 2A1 1 0 0 1 18 9V8H7a3 3 0 0 0-3 3a1 1 0 1 1-2 0a5 5 0 0 1 5-5h11V5ZM6 19a1 1 0 0 1-1.707.707l-2-2a1 1 0 0 1 0-1.414l2-2A1 1 0 0 1 6 15v1h11a3 3 0 0 0 3-3a1 1 0 1 1 2 0a5 5 0 0 1-4.998 5H6v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}