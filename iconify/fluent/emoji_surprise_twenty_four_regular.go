package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiSurpriseTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2.004c5.523 0 10 4.477 10 10s-4.477 10-10 10s-10-4.477-10-10s4.477-10 10-10Zm0 1.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17Zm0 9.495a2.25 2.25 0 1 1 0 4.5a2.25 2.25 0 0 1 0-4.5ZM9 8.75a1.25 1.25 0 1 1 0 2.499A1.25 1.25 0 0 1 9 8.75Zm6 0a1.25 1.25 0 1 1 0 2.499a1.25 1.25 0 0 1 0-2.499Z"/>`),
		g.Group(children),
	)
}