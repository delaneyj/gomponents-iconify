package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.645 20.5a3.502 3.502 0 0 0 6.71 0h-6.71ZM3 19.5h18v-3l-2-3v-5a7 7 0 1 0-14 0v5l-2 3v3Z"/>`),
		g.Group(children),
	)
}