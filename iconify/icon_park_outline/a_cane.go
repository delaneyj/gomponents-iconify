package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ACane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19.558 44.768L33.642 18.28c1.173-2.207 3.811-9.299-3.252-13.055C23.326 1.47 19.157 7.181 17.749 9.83"/>`),
		g.Group(children),
	)
}