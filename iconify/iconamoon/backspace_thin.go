package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackspaceThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.92 5a2 2 0 0 0-1.519.698l-4.285 5a2 2 0 0 0 0 2.604l4.285 5A2 2 0 0 0 7.92 19H19a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H7.92ZM15 10l-4 4m0-4l4 4"/>`),
		g.Group(children),
	)
}