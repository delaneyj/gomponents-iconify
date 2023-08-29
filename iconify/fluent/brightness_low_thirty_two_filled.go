package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessLowThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 5a1 1 0 0 0-2 0v2a1 1 0 1 0 2 0V5Zm6 11a7 7 0 1 1-14 0a7 7 0 0 1 14 0Zm-2 0a5 5 0 0 0-5-5v10a5 5 0 0 0 5-5Zm7 0a1 1 0 0 1-1 1h-2a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1Zm-11 9a1 1 0 0 0-2 0v2a1 1 0 1 0 2 0v-2Zm-9-9a1 1 0 0 1-1 1H5a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1Zm.707-8.707a1 1 0 0 0-1.414 1.414l1 1a1 1 0 0 0 1.414-1.414l-1-1ZM7.293 24.707a1 1 0 0 0 1.414 0l1-1a1 1 0 1 0-1.414-1.414l-1 1a1 1 0 0 0 0 1.414Zm16-17.414a1 1 0 0 1 1.414 1.414l-1 1a1 1 0 0 1-1.414-1.414l1-1Zm0 17.414l-1-1a1 1 0 0 1 1.414-1.414l1 1a1 1 0 0 1-1.414 1.414Z"/>`),
		g.Group(children),
	)
}