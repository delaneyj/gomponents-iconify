package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrthodoxCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M55 32.001v-7.5H35.68V17h5.521V9.5H35.68V2h-7.36v7.5h-5.519V17h5.519v7.501H9v7.5h19.32v13.437l-7.359-3.125v7.5l7.359 3.125V62h7.36v-5.937l7.361 3.125v-7.5l-7.361-3.125V32.001z"/>`),
		g.Group(children),
	)
}