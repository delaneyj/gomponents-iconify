package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M63.915 57.025a6.89 6.89 0 0 1-6.887 6.895H6.886A6.89 6.89 0 0 1 0 57.025V6.891A6.889 6.889 0 0 1 6.886 0h50.142a6.889 6.889 0 0 1 6.887 6.891v50.134z"/><path fill="#fff" d="m25.2 17.929l24.308-2.201l-1.269 24.292c-1.924 1.712-4.559 1.971-6.07.52l-2.568-2.462l-11.553 12.04a5.092 5.092 0 0 1-7.193.151l-4.759-4.57a5.088 5.088 0 0 1-.149-7.193l11.554-12.04l-2.566-2.46c-1.508-1.449-1.363-4.084.261-6.08"/>`),
		g.Group(children),
	)
}