package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 1h6v8.5h6V23H3V9.5h6V1Zm2 2v8.5H5V14h14v-2.5h-6V3h-2Zm8 13H5v5h9v-3h2v3h3v-5Z"/>`),
		g.Group(children),
	)
}