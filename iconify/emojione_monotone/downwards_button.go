package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownwardsButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.212 10.785c-11.715-11.714-30.711-11.714-42.426 0c-11.715 11.717-11.715 30.711 0 42.427c11.715 11.717 30.711 11.717 42.426 0c11.717-11.716 11.717-30.71 0-42.427M31.999 47l-20-20h40l-20 20"/>`),
		g.Group(children),
	)
}