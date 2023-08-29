package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationOffFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 16.586V13a9.883 9.883 0 0 0-1.087-4.499L30 3.414L28.586 2L2 28.586L3.414 30l6-6H11v1a5 5 0 0 0 10 0v-1h7a1 1 0 0 0 1-1v-3a1 1 0 0 0-.293-.707zM19 25a3 3 0 0 1-6 0v-1h6zm2.61-20.262A9.926 9.926 0 0 0 17 3.051V1h-2v2.05A10.014 10.014 0 0 0 6 13v3.586l-2.707 2.707A1 1 0 0 0 3 20v3a.974.974 0 0 0 .059.29z"/>`),
		g.Group(children),
	)
}