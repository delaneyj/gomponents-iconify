package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionRotateNinetyLeftTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m20.385 15.327l-9-3.75a1 1 0 1 0-.77 1.846l2.385.993v3.667l-2.385.994a1 1 0 0 0 .77 1.846l9-3.75a1 1 0 0 0 0-1.846ZM15 15.25l2.4 1l-2.4 1v-2ZM8 20a1 1 0 1 1-2 0V6.414l-.293.293a1 1 0 0 1-1.414-1.414l2-2a1 1 0 0 1 1.414 0l2 2a1 1 0 0 1-1.414 1.414L8 6.414V20Zm8-9a1 1 0 0 1-1-1V6.414l-.293.293a1 1 0 1 1-1.414-1.414l2-2a1 1 0 0 1 1.414 0l2 2a1 1 0 0 1-1.414 1.414L17 6.414V10a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}