package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriveTrainTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 2a2 2 0 0 0-2 2v1h-1.585a1.5 1.5 0 0 0-2.83 0H7V4a2 2 0 1 0-4 0v3a2 2 0 1 0 4 0V6h1.585c.151.426.489.764.915.915v6.17a1.508 1.508 0 0 0-.915.915H7v-1a2 2 0 1 0-4 0v3a2 2 0 1 0 4 0v-1h1.585a1.5 1.5 0 0 0 2.83 0H13v1a2 2 0 1 0 4 0v-3a2 2 0 1 0-4 0v1h-1.585a1.504 1.504 0 0 0-.915-.915v-6.17c.426-.151.764-.489.915-.915H13v1a2 2 0 1 0 4 0V4a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}