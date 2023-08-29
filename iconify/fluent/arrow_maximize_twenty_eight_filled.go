package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMaximizeTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.5 3H24a1 1 0 0 1 .993.883L25 4v7.5a1 1 0 0 1-1.993.117L23 11.5V6.414L6.414 23H11.5a1 1 0 0 1 .993.883L12.5 24a1 1 0 0 1-.883.993L11.5 25H4a1 1 0 0 1-.993-.883L3 24v-7.5a1 1 0 0 1 1.993-.117L5 16.5v5.085L21.585 5H16.5a1 1 0 0 1-.993-.883L15.5 4a1 1 0 0 1 .883-.993L16.5 3H24h-7.5Z"/>`),
		g.Group(children),
	)
}