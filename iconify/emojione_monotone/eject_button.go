package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.214 10.787c-11.715-11.716-30.711-11.716-42.426 0c-11.717 11.715-11.717 30.711 0 42.426c11.715 11.716 30.711 11.716 42.426 0c11.715-11.715 11.715-30.711 0-42.426M48 43.351H16V37.65h32v5.701m-32-9.702l16-19l16 19H16"/>`),
		g.Group(children),
	)
}