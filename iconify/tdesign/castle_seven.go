package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 1v1.083c2.838.476 5 2.944 5 5.917h4v14H2V8h4a6.002 6.002 0 0 1 5-5.917V1h2ZM8 8h8a4 4 0 0 0-8 0Zm-4 4.535A3.982 3.982 0 0 1 6 12a3.99 3.99 0 0 1 3 1.354A3.99 3.99 0 0 1 12 12a3.99 3.99 0 0 1 3 1.354A3.99 3.99 0 0 1 18 12c.729 0 1.412.195 2 .535V10H4v2.535ZM8 16a2 2 0 1 0-4 0v4h4v-4Zm2 4h4v-4a2 2 0 1 0-4 0v4Zm6 0h4v-4a2 2 0 1 0-4 0v4Z"/>`),
		g.Group(children),
	)
}