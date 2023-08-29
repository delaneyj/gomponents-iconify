package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crehana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.371 0 0 5.371 0 12c0 6.626 5.371 12 12 12s12-5.374 12-12c0-6.629-5.374-12-12-12zm5.94 9.843v7.915h-3.957v-3.892h-3.895v3.83H6.13v-3.957h3.833V9.843H6.06V5.948h3.957v3.895h3.965V5.948h3.957v3.895z"/>`),
		g.Group(children),
	)
}