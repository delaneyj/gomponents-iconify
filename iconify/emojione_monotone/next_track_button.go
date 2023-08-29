package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextTrackButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.213 10.788c-11.715-11.717-30.711-11.717-42.426 0c-11.716 11.715-11.716 30.711 0 42.426c11.715 11.715 30.711 11.715 42.426 0c11.716-11.715 11.716-30.711 0-42.426zM52 46h-5.219V32l-17.39 14V32L12 46V18l17.392 14V18l17.39 14V18H52v28z"/>`),
		g.Group(children),
	)
}