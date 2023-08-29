package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 21h-5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h5v2h-5v6h5zm10 0h-5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h5v2h-5v6h5z"/><path fill="currentColor" d="M28 26H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h24a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2ZM4 8v16h24V8Z"/>`),
		g.Group(children),
	)
}