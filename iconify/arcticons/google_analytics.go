package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAnalytics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.5 42.5h0a5 5 0 0 0 5-5v-27a5 5 0 0 0-5-5h0a5 5 0 0 0-5 5v27a5 5 0 0 0 5 5Zm-13.5 0h0a5 5 0 0 0 5-5v-16a5 5 0 0 0-5-5h0a5 5 0 0 0-5 5v16a5 5 0 0 0 5 5Z"/><circle cx="10.5" cy="37.5" r="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}