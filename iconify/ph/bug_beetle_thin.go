package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugBeetleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 140h20a4 4 0 0 0 0-8h-20v-24h20a4 4 0 0 0 0-8h-20.11a75.8 75.8 0 0 0-23.37-50.86l22.31-22.31a4 4 0 1 0-5.66-5.66L174.44 43.9a75.77 75.77 0 0 0-92.88 0L58.83 21.17a4 4 0 0 0-5.66 5.66l22.31 22.31A75.8 75.8 0 0 0 52.11 100H32a4 4 0 0 0 0 8h20v24H32a4 4 0 0 0 0 8h20v12a75.41 75.41 0 0 0 1 12H32a4 4 0 0 0 0 8h22.69a76 76 0 0 0 146.62 0H224a4 4 0 0 0 0-8h-21a75.41 75.41 0 0 0 1-12ZM128 36a68.08 68.08 0 0 1 67.87 64H60.13A68.08 68.08 0 0 1 128 36Zm4 183.87V136a4 4 0 0 0-8 0v83.87A68.08 68.08 0 0 1 60 152v-44h136v44a68.08 68.08 0 0 1-64 67.87Z"/>`),
		g.Group(children),
	)
}