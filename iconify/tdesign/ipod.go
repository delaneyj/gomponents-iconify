package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h18v22H3V1Zm2 2v10h2V9h2v4h2V5h2v8h2V7h2v6h2V3H5Zm14 12H5v6h14v-6Zm-9 3a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/>`),
		g.Group(children),
	)
}