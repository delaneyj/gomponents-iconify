package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteSmallSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#d0d2d3" d="M43.36 39.932a2.18 2.18 0 0 1-2.181 2.181H25.316a2.182 2.182 0 0 1-2.181-2.181V24.068c0-1.205.978-2.181 2.181-2.181h15.863a2.18 2.18 0 0 1 2.181 2.181v15.864"/>`),
		g.Group(children),
	)
}