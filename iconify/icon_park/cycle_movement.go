package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CycleMovement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M20 15L24 19L28 15"/><path d="M24 19V8C24 5.79086 25.7909 4 28 4H40C42.2091 4 44 5.79086 44 8V19"/><path d="M28 33L24 29L20 33"/><path d="M24 29V40C24 42.2091 22.2091 44 20 44H8C5.79086 44 4 42.2091 4 40V29"/><path d="M33 20L29 24L33 28"/><path d="M29 24H40C42.2091 24 44 25.7909 44 28V40C44 42.2091 42.2091 44 40 44H29"/><path d="M15 28L19 24L15 20"/><path d="M19 24H8C5.79086 24 4 22.2091 4 20V8C4 5.79086 5.79086 4 8 4H19"/></g>`),
		g.Group(children),
	)
}