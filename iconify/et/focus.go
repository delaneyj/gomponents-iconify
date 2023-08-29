package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Focus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8.5 0h-7C.673 0 0 .673 0 1.5v6a.5.5 0 0 0 1 0v-6a.5.5 0 0 1 .5-.5h7a.5.5 0 0 0 0-1zm32 0h-7a.5.5 0 0 0 0 1h7a.5.5 0 0 1 .5.5v6a.5.5 0 0 0 1 0v-6c0-.827-.673-1.5-1.5-1.5zm1 24a.5.5 0 0 0-.5.5v6a.5.5 0 0 1-.5.5h-7a.5.5 0 0 0 0 1h7c.827 0 1.5-.673 1.5-1.5v-6a.5.5 0 0 0-.5-.5zm-41 0a.5.5 0 0 0-.5.5v6c0 .827.673 1.5 1.5 1.5h7a.5.5 0 0 0 0-1h-7a.5.5 0 0 1-.5-.5v-6a.5.5 0 0 0-.5-.5zm23-5a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5h-5a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5h5zM19 14h4v4h-4v-4z"/>`),
		g.Group(children),
	)
}