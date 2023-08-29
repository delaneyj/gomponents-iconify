package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightCircleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2zm3.25 6a.75.75 0 0 0-.75.75v4.639L9.278 8.217a.75.75 0 0 0-1.056 1.066L13.49 14.5H8.75a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 .75-.75v-6.5a.75.75 0 0 0-.75-.75z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}