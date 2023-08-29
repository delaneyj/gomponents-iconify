package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpwardsButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.213 10.788c-11.715-11.717-30.711-11.717-42.426 0c-11.716 11.715-11.716 30.71 0 42.426c11.715 11.715 30.711 11.715 42.426 0c11.717-11.716 11.717-30.711 0-42.426zM12 38l20-20l20 20H12z"/>`),
		g.Group(children),
	)
}