package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastDownButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.213 10.787c-11.715-11.715-30.711-11.715-42.426 0c-11.716 11.717-11.716 30.711 0 42.426c11.715 11.715 30.711 11.715 42.426 0c11.716-11.715 11.716-30.71 0-42.426zM50 32.723L32 51.998L14 32.723h11.879L14 20h36L38.121 32.723H50z"/>`),
		g.Group(children),
	)
}