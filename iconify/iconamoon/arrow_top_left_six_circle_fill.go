package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopLeftSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm8-3a1 1 0 0 0-1 1v4a1 1 0 1 0 2 0v-3h3a1 1 0 1 0 0-2h-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}