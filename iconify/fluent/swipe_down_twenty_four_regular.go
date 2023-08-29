package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeDownTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 5.997a.75.75 0 0 1 .743.647l.007.102V19.44l2.22-2.216a.75.75 0 0 1 .976-.073l.084.073c.267.266.29.682.073.975l-.073.084l-3.5 3.497a.75.75 0 0 1-.976.073l-.084-.073l-3.5-3.497a.749.749 0 0 1 .976-1.132l.084.073l2.22 2.216V6.746a.75.75 0 0 1 .75-.75ZM12 2c2.761 0 5 2.237 5 4.996a4.998 4.998 0 0 1-3.25 4.68v-1.651A3.499 3.499 0 0 0 12 3.499a3.499 3.499 0 0 0-1.75 6.526v1.652A4.998 4.998 0 0 1 7 6.996A4.998 4.998 0 0 1 12 2Z"/>`),
		g.Group(children),
	)
}