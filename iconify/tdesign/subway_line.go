package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubwayLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2v3h9.17A3.008 3.008 0 0 1 18 3.17V2h2v1.17a3.001 3.001 0 0 1 0 5.66V19H8.83A3.008 3.008 0 0 1 7 20.83V22H5v-1.17A3.008 3.008 0 0 1 3.17 19H2v-2h1.17A3.008 3.008 0 0 1 5 15.17V7H2V5h3V2h2Zm0 5v8.17A3.008 3.008 0 0 1 8.83 17H18V8.83A3.008 3.008 0 0 1 16.17 7H7Zm12-2a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM6 17a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}