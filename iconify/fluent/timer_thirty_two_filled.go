package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 2a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2h-6Zm15 16c0 6.627-5.373 12-12 12S4 24.627 4 18S9.373 6 16 6s12 5.373 12 12Zm-11-6a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0v-6Zm8.293-5.707a1 1 0 0 1 1.414 0l2 2a1 1 0 1 1-1.414 1.414l-2-2a1 1 0 0 1 0-1.414Z"/>`),
		g.Group(children),
	)
}