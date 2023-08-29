package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h10v-2.5l3.5 3.5l-3.5 3.5V19H7v2.5L3.5 18L7 14.5V17Zm6-11v9h-2V6H5V4h14v2h-6Z"/>`),
		g.Group(children),
	)
}