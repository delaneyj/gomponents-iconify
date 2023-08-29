package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.929.515L21.07 14.657l-1.414 1.414l-3.823-3.822L14 14.999v7h-4v-7L4 6H3V4h4.585l-2.07-2.07L6.929.515ZM21 4v2h-1l-1.915 2.872L13.213 4H21Z"/>`),
		g.Group(children),
	)
}