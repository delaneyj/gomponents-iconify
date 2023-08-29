package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagicWandBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M252 152a12 12 0 0 1-12 12h-12v12a12 12 0 0 1-24 0v-12h-12a12 12 0 0 1 0-24h12v-12a12 12 0 0 1 24 0v12h12a12 12 0 0 1 12 12ZM56 76h12v12a12 12 0 0 0 24 0V76h12a12 12 0 1 0 0-24H92V40a12 12 0 0 0-24 0v12H56a12 12 0 0 0 0 24Zm128 112h-4v-4a12 12 0 0 0-24 0v4h-4a12 12 0 0 0 0 24h4v4a12 12 0 0 0 24 0v-4h4a12 12 0 0 0 0-24Zm38.14-105.17L82.82 222.14a20 20 0 0 1-28.28 0l-20.69-20.68a20 20 0 0 1 0-28.29L173.17 33.86a20 20 0 0 1 28.28 0l20.69 20.68a20 20 0 0 1 0 28.29ZM159 112l-15-15l-90.35 90.31l15 15Zm43.31-43.31l-15-15L161 80l15 15Z"/>`),
		g.Group(children),
	)
}