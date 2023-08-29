package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TouchOneDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 24h-2a5 5 0 0 1-10 0H6a7 7 0 0 0 14 0Z"/><path fill="currentColor" d="M28 14V8a7.008 7.008 0 0 0-7-7h-5a6.146 6.146 0 0 0-4.106 1.566l-8.01 7.308a2.999 2.999 0 0 0 3.88 4.55l.001.001L10 12.895V24a3 3 0 0 0 6 0v-5.184a2.939 2.939 0 0 0 3.53-1.217A2.963 2.963 0 0 0 21 18a2.994 2.994 0 0 0 2.53-1.401A2.963 2.963 0 0 0 25 17a3.003 3.003 0 0 0 3-3Zm-2 0a1 1 0 0 1-2 0v-1h-2v2a1 1 0 0 1-2 0v-2h-2v3a1 1 0 0 1-2 0v-3h-2v11a1 1 0 0 1-2 0V9.104l-5.4 3.697a1 1 0 0 1-1.308-1.505l7.95-7.251A4.148 4.148 0 0 1 16 3h5a5.006 5.006 0 0 1 5 5Z"/>`),
		g.Group(children),
	)
}