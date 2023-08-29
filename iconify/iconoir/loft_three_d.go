package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoftThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 17c-9 0-11 6-20 6M22 1C13 1 11 7 2 7m10 9.5v-9m0 9l2.5-2.5M12 16.5L9.5 14M12 7.5l2.5 2.5M12 7.5L9.5 10"/>`),
		g.Group(children),
	)
}