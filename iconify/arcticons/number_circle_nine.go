package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M24.438 36.15c4.407-.23 7.712-4.186 7.712-8.6V20A8.15 8.15 0 0 0 24 11.85h0A8.15 8.15 0 0 0 15.85 20"/><path d="M15.85 20A8.15 8.15 0 0 0 24 28.148h0a8.15 8.15 0 0 0 8.15-8.15M17.282 33.487c1.49 1.94 3.358 2.663 5.958 2.663h1.198"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}