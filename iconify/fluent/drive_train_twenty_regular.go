package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriveTrainTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4a2 2 0 1 1 4 0v1h1.585a1.5 1.5 0 0 1 2.83 0H13V4a2 2 0 1 1 4 0v3a2 2 0 1 1-4 0V6h-1.585a1.508 1.508 0 0 1-.915.915v6.17c.426.151.764.489.915.915H13v-1a2 2 0 1 1 4 0v3a2 2 0 1 1-4 0v-1h-1.585a1.5 1.5 0 0 1-2.83 0H7v1a2 2 0 1 1-4 0v-3a2 2 0 1 1 4 0v1h1.585c.151-.426.489-.764.915-.915v-6.17A1.504 1.504 0 0 1 8.585 6H7v1a2 2 0 1 1-4 0V4Zm2-1a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0V4a1 1 0 0 0-1-1Zm10 0a1 1 0 0 0-1 1v3a1 1 0 1 0 2 0V4a1 1 0 0 0-1-1ZM5 12a1 1 0 0 0-1 1v3a1 1 0 1 0 2 0v-3a1 1 0 0 0-1-1Zm9 1v3a1 1 0 1 0 2 0v-3a1 1 0 1 0-2 0Z"/>`),
		g.Group(children),
	)
}