package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204.73 51.85A108.07 108.07 0 0 0 20 128v56a28 28 0 0 0 28 28h16a28 28 0 0 0 28-28v-40a28 28 0 0 0-28-28H44.84A84.05 84.05 0 0 1 128 44h.64a83.7 83.7 0 0 1 82.52 72H192a28 28 0 0 0-28 28v40a28 28 0 0 0 28 28h19.6a20 20 0 0 1-19.6 16h-56a12 12 0 0 0 0 24h56a44.05 44.05 0 0 0 44-44v-80a107.34 107.34 0 0 0-31.27-76.15ZM64 140a4 4 0 0 1 4 4v40a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4v-44Zm124 44v-40a4 4 0 0 1 4-4h20v48h-20a4 4 0 0 1-4-4Z"/>`),
		g.Group(children),
	)
}