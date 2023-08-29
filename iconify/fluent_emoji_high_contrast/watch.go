package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M16 8a1 1 0 0 0-1 1v5.268a2 2 0 0 0-.864 2.46l-2.226 2.481a1 1 0 0 0 1.49 1.336l2.301-2.567A2 2 0 0 0 17 14.268V9a1 1 0 0 0-1-1Z"/><path d="M9 3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3.252A12 12 0 0 1 27.834 14h.666a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-.666A12 12 0 0 1 23 25.748V29a1 1 0 0 1-1 1H10a1 1 0 0 1-1-1v-3.252A11.984 11.984 0 0 1 4 16c0-4.015 1.972-7.57 5-9.748V3Zm16 13a9 9 0 1 0-18 0a9 9 0 0 0 18 0Z"/></g>`),
		g.Group(children),
	)
}