package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircleTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 12a8 8 0 0 1 8-8a1 1 0 1 0 0-2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10a1 1 0 1 0-2 0a8 8 0 1 1-16 0Zm15.707-6.293a1 1 0 0 0-1.414-1.414L16 6.586l-.293-.293a1 1 0 1 0-1.414 1.414l1 1a1 1 0 0 0 1.414 0l3-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}