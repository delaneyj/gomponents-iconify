package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackSmallSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#354a54" d="M42.611 39.31c0 1.2-.972 2.173-2.171 2.173H24.643a2.172 2.172 0 0 1-2.17-2.173V23.515a2.17 2.17 0 0 1 2.17-2.171H40.44c1.199 0 2.171.973 2.171 2.171V39.31"/>`),
		g.Group(children),
	)
}