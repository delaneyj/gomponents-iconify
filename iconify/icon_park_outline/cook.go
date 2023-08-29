package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 36V20H14v16a6 6 0 0 0 6 6h16a6 6 0 0 0 6-6ZM4 20h40M18 8v4m10-6v6m10-4v4"/>`),
		g.Group(children),
	)
}