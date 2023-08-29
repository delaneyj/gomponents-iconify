package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.213 10.786c-11.715-11.715-30.711-11.715-42.426 0c-11.716 11.715-11.716 30.711 0 42.426c11.715 11.717 30.711 11.717 42.426 0c11.716-11.715 11.716-30.711 0-42.426M36.821 53.999l-8.179-8.912h4.596V22.884c0-3.045-2.41-5.521-5.373-5.521a5.27 5.27 0 0 0-3.8 1.617L19 13.771C21.368 11.339 24.517 10 27.865 10c6.914 0 12.537 5.779 12.537 12.885v22.203H45l-8.179 8.911"/>`),
		g.Group(children),
	)
}