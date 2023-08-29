package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smoking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M37 11a1 1 0 0 1 1-1a5 5 0 0 1 5 5v1.59a4.4 4.4 0 0 1-1.644 3.436A5.994 5.994 0 0 1 44 25v3a1 1 0 1 1-2 0v-3a4 4 0 0 0-4-4a1 1 0 1 1 0-2h.59A2.41 2.41 0 0 0 41 16.59V15a3 3 0 0 0-3-3a1 1 0 0 1-1-1Z"/><path d="M34 17a3 3 0 0 0-3 3v1.818a3 3 0 0 0 3 3h3a3 3 0 0 1 3 3V28a1 1 0 1 1-2 0v-.182a1 1 0 0 0-1-1h-3a5 5 0 0 1-5-5V20a5 5 0 0 1 5-5a1 1 0 1 1 0 2ZM7 31a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2H7Zm33 1a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4Zm4 0a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4Z"/></g>`),
		g.Group(children),
	)
}