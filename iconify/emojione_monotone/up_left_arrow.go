package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.212 10.788c-11.715-11.717-30.711-11.717-42.426 0c-11.715 11.715-11.715 30.708 0 42.425c11.715 11.715 30.711 11.715 42.426 0c11.717-11.717 11.717-30.711 0-42.425m-14.25 35.21L25.509 32.717L18 40.131V17.999h22.057l-7.432 7.564L46 38.953l-7.038 7.045"/>`),
		g.Group(children),
	)
}