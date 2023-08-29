package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastTrackButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M10.787 53.213c11.715 11.717 30.711 11.717 42.426 0c11.716-11.715 11.716-30.711 0-42.426c-11.715-11.715-30.711-11.715-42.426 0c-11.716 11.714-11.716 30.711 0 42.426zM12 18h5.219v14l17.39-14v14L52 18v28L34.608 32v14l-17.39-14v14H12V18z"/>`),
		g.Group(children),
	)
}