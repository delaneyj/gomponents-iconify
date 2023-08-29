package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M63.794 56.915a6.877 6.877 0 0 1-6.878 6.882H6.874A6.878 6.878 0 0 1 0 56.915V6.878A6.877 6.877 0 0 1 6.874 0h50.042a6.876 6.876 0 0 1 6.878 6.878v50.037z"/><path fill="#fff" d="m39.19 45.666l-24.26 2.196l1.264-24.25c1.922-1.708 4.55-1.967 6.06-.519l2.561 2.457l11.532-12.02a5.082 5.082 0 0 1 7.179-.151l4.749 4.562a5.078 5.078 0 0 1 .149 7.18l-11.53 12.02l2.559 2.456c1.505 1.445 1.36 4.075-.263 6.068"/>`),
		g.Group(children),
	)
}