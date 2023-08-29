package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberThreeCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm8-6a1 1 0 0 0 0 2h3l-1.8 2.4A1 1 0 0 0 12 12a2 2 0 1 1-1.333 3.491a1 1 0 1 0-1.334 1.49a4 4 0 1 0 4.379-6.597L15.8 7.6A1 1 0 0 0 15 6h-5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}