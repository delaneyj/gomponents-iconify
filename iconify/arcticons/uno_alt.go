package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnoAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="13.915" ry="8.293" transform="rotate(-73.087 24 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.198 37.381l7.604-26.762M15.659 24h16.699"/><rect width="24" height="39" x="12" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}