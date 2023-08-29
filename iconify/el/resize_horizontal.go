package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M304.102 295.898L0 600l304.102 304.102v-203.54h591.797v203.539L1200 600L895.898 295.898v203.539H304.102V295.898z"/>`),
		g.Group(children),
	)
}