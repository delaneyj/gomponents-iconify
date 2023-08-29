package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBarVerticalThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 7a3 3 0 1 1 6 0v18a3 3 0 1 1-6 0V7Zm3-1a1 1 0 0 0-1 1v18a1 1 0 1 0 2 0V7a1 1 0 0 0-1-1Zm5 7a3 3 0 1 1 6 0v12a3 3 0 1 1-6 0V13Zm3-1a1 1 0 0 0-1 1v12a1 1 0 1 0 2 0V13a1 1 0 0 0-1-1Zm8 4a3 3 0 0 0-3 3v6a3 3 0 1 0 6 0v-6a3 3 0 0 0-3-3Zm-1 3a1 1 0 1 1 2 0v6a1 1 0 1 1-2 0v-6Z"/>`),
		g.Group(children),
	)
}