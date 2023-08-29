package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraDollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 3.998h17v4.434l7-4.2v15.49l-7-4v4.276H0v-16Zm17 9.42l5 2.857v-8.51l-5 3v2.653ZM2 5.998v12h13v-12H2ZM9.5 7v1h2v2h-4v1h2a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2v1h-2v-1h-2v-2h4v-1h-2a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2V7h2Z"/>`),
		g.Group(children),
	)
}