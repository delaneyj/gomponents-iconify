package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryCharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.376 5.172L8.79 11h6l-5.462 8.876l-1.704-1.048L11.21 13h-6l5.462-8.876l1.704 1.048ZM0 5h7.5v2H2v10h4v2H0V5Zm14 0h7v14h-8v-2h6V7h-5V5Zm10 4v6h-2V9h2Z"/>`),
		g.Group(children),
	)
}