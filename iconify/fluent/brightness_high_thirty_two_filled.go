package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessHighThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 3a1 1 0 0 0-2 0v2a1 1 0 1 0 2 0V3Zm6 13a7 7 0 1 1-14 0a7 7 0 0 1 14 0Zm-2 0a5 5 0 0 0-5-5v10a5 5 0 0 0 5-5Zm9 0a1 1 0 0 1-1 1h-2a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1ZM17 27a1 1 0 0 0-2 0v2a1 1 0 1 0 2 0v-2ZM6 16a1 1 0 0 1-1 1H3a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1Zm.707-10.706a1 1 0 1 0-1.414 1.414l2 2a1 1 0 0 0 1.414-1.414l-2-2ZM5.293 26.708a1 1 0 0 0 1.414 0l2-2a1 1 0 1 0-1.414-1.414l-2 2a1 1 0 0 0 0 1.414Zm20-21.414a1 1 0 0 1 1.414 1.414l-2 2a1 1 0 0 1-1.414-1.414l2-2Zm0 21.414l-2-2a1 1 0 0 1 1.414-1.414l2 2a1 1 0 0 1-1.414 1.414Z"/>`),
		g.Group(children),
	)
}