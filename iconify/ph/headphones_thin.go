package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M199.05 65.48A100.07 100.07 0 0 0 28 136v56a20 20 0 0 0 20 20h16a20 20 0 0 0 20-20v-40a20 20 0 0 0-20-20H36.08A92 92 0 0 1 128 44h.7a91.75 91.75 0 0 1 91.22 88H192a20 20 0 0 0-20 20v40a20 20 0 0 0 20 20h16a20 20 0 0 0 20-20v-56a99.43 99.43 0 0 0-28.95-70.52ZM64 140a12 12 0 0 1 12 12v40a12 12 0 0 1-12 12H48a12 12 0 0 1-12-12v-52Zm156 52a12 12 0 0 1-12 12h-16a12 12 0 0 1-12-12v-40a12 12 0 0 1 12-12h28Z"/>`),
		g.Group(children),
	)
}