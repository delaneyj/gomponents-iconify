package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.213 10.786c-11.715-11.715-30.711-11.715-42.426 0c-11.716 11.717-11.716 30.711 0 42.426c11.715 11.717 30.711 11.717 42.426 0c11.716-11.715 11.716-30.709 0-42.426M41.056 44.999H19V22.866l7.51 7.416l13.451-13.283L47 24.044L33.626 37.437l7.43 7.562"/>`),
		g.Group(children),
	)
}