package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 4h-3V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v1H3a1 1 0 0 0-1 1v3a4 4 0 0 0 4 4h1.54A6 6 0 0 0 11 13.91V16h-1a3 3 0 0 0-3 3v2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-2a3 3 0 0 0-3-3h-1v-2.09A6 6 0 0 0 16.46 12H18a4 4 0 0 0 4-4V5a1 1 0 0 0-1-1ZM6 10a2 2 0 0 1-2-2V6h2v2a6 6 0 0 0 .35 2Zm8 8a1 1 0 0 1 1 1v1H9v-1a1 1 0 0 1 1-1Zm2-10a4 4 0 0 1-8 0V4h8Zm4 0a2 2 0 0 1-2 2h-.35A6 6 0 0 0 18 8V6h2Z"/>`),
		g.Group(children),
	)
}