package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteControlFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 12a1 1 0 0 1 1 1v9H6v-9a1 1 0 0 1 1-1h10Zm-7 2H8v2h2v-2Zm2-8a6 6 0 0 1 5.368 3.316l-1.79.895a4 4 0 0 0-7.156 0l-1.79-.895A6 6 0 0 1 12 6Zm0-4a10 10 0 0 1 8.947 5.527l-1.79.895A8 8 0 0 0 12 4a8 8 0 0 0-7.157 4.422l-1.79-.895A10 10 0 0 1 12 2Z"/>`),
		g.Group(children),
	)
}