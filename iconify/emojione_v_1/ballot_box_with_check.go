package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallotBoxWithCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#405866" d="M64 57.1a6.895 6.895 0 0 1-6.897 6.898H6.893A6.897 6.897 0 0 1-.004 57.1V6.9A6.9 6.9 0 0 1 6.893.002h50.2A6.896 6.896 0 0 1 63.99 6.9v50.2"/><path fill="#f4f4f4" d="M54.04 14.186a4.668 4.668 0 0 0-6.582.516l-20.1 23.515l-10.815-9.03a4.27 4.27 0 0 0-6.01.538a4.258 4.258 0 0 0 .54 6l14.492 12.1a4.238 4.238 0 0 0 3.248.949a4.663 4.663 0 0 0 3.189-1.622l22.554-26.387a4.668 4.668 0 0 0-.516-6.581"/>`),
		g.Group(children),
	)
}