package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudArrowUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M252 128a91.18 91.18 0 0 1-18.41 55.21a12 12 0 0 1-19.18-14.42A68 68 0 1 0 92 128a12 12 0 0 1-24 0a91.7 91.7 0 0 1 2.19-20A44 44 0 0 0 72 196h24a12 12 0 0 1 0 24H72a68 68 0 1 1 7-135.63A92 92 0 0 1 252 128Zm-91.51-8.49a12 12 0 0 0-17 0l-32 32a12 12 0 1 0 17 17L140 157v51a12 12 0 0 0 24 0v-51l11.51 11.52a12 12 0 0 0 17-17Z"/>`),
		g.Group(children),
	)
}