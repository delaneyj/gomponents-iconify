package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M63.791 56.913a6.877 6.877 0 0 1-6.876 6.882H6.874A6.878 6.878 0 0 1 0 56.913V6.877A6.876 6.876 0 0 1 6.874 0h50.042c3.8 0 6.876 3.08 6.876 6.877v50.036z"/><path fill="#fff" d="M47.865 35.77L31.976 54.233L16.008 35.947c.19-2.566 1.897-4.581 3.988-4.591l3.548-.018l-.09-16.659a5.083 5.083 0 0 1 5.05-5.104l6.586-.033a5.08 5.08 0 0 1 5.102 5.05l.087 16.655l3.546-.02c2.088-.006 3.815 1.98 4.04 4.542"/>`),
		g.Group(children),
	)
}