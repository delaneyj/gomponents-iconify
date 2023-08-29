package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalPrivacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="11" r="7"/><path d="M4 41c0-8.837 8.059-16 18-16m5 6h14v10H27zm10 0v-3a3 3 0 1 0-6 0v3"/></g>`),
		g.Group(children),
	)
}