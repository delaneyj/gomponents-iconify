package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowflake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 4V44"/><path d="M6.72461 14L41.3656 34"/><path d="M6.71923 33.9773L41.2814 14.0228"/><path d="M12 10L15 19L6 21"/><path d="M6 27L15 29L12 38"/><path d="M36 10L33 19L42 21"/><path d="M42 27L33 29L36 38"/><path d="M18 7L24 13L30 7"/><path d="M18 41L24 35L30 41"/></g>`),
		g.Group(children),
	)
}