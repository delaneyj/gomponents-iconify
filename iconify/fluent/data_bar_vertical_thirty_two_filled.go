package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBarVerticalThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 7a3 3 0 1 1 6 0v18a3 3 0 1 1-6 0V7Zm8 6a3 3 0 1 1 6 0v12a3 3 0 1 1-6 0V13Zm11 3a3 3 0 0 0-3 3v6a3 3 0 1 0 6 0v-6a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}