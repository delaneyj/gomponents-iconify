package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteMediumSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#d0d2d3" d="M53.935 49.29a4.71 4.71 0 0 1-4.707 4.713h-34.27a4.71 4.71 0 0 1-4.708-4.713V15.03a4.71 4.71 0 0 1 4.708-4.709h34.27c2.6 0 4.707 2.109 4.707 4.709v34.26"/>`),
		g.Group(children),
	)
}