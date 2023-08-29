package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryAlertVariantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20H6V6h8m.67-2H13V2H7v2H5.33C4.6 4 4 4.6 4 5.33v15.34C4 21.4 4.6 22 5.33 22h9.34c.73 0 1.33-.6 1.33-1.33V5.33C16 4.6 15.4 4 14.67 4M21 7h-2v6h2V8m0 7h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}