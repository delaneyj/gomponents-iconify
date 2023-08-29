package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubscriptTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 6v13H9V6H3V4h14v2h-6Zm8.55 10.58a.8.8 0 1 0-1.32-.36l-1.154.33A2.001 2.001 0 1 1 21 16c0 .573-.24 1.09-.627 1.454L18.744 19H21v1h-4v-1l2.55-2.42Z"/>`),
		g.Group(children),
	)
}