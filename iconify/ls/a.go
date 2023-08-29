package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func A(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M491 222v-94h72v526h-72v-74c-50 55-124 91-206 91C133 671 0 551 0 401s133-270 285-270c82 0 156 36 206 91zm0 185v-12c-4-110-95-198-207-198c-115 0-215 91-215 204s100 204 215 204c112 0 203-88 207-198z"/>`),
		g.Group(children),
	)
}