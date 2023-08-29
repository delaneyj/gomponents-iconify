package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4h8v6h-8V4m0 17V11h8v10h-8m-9 0v-6h8v6H3m0-7V4h8v10H3m1-9v8h6V5H4m9 0v4h6V5h-6m0 7v8h6v-8h-6m-9 4v4h6v-4H4Z"/>`),
		g.Group(children),
	)
}