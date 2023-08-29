package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastReverseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.212 10.787c-11.715-11.715-30.711-11.715-42.426 0c-11.715 11.717-11.715 30.712 0 42.427s30.711 11.715 42.426 0c11.717-11.715 11.717-30.71 0-42.427m-8.213 39.212L32.276 38.122v11.877L12.999 32.001L32.276 14v11.879L44.999 14v35.999"/>`),
		g.Group(children),
	)
}