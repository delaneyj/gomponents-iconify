package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastUpButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.214 10.787c-11.715-11.715-30.711-11.715-42.426 0c-11.717 11.717-11.717 30.711 0 42.426c11.715 11.715 30.711 11.715 42.426 0s11.715-30.71 0-42.426m-3.213 35.211h-36L25.88 33.279H14.001l18-19.28l18 19.28H38.122l11.879 12.719"/>`),
		g.Group(children),
	)
}