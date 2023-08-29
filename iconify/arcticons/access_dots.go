package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="15.046" cy="24" r="6.546" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.954" cy="24" r="6.546" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.501 24a19.406 19.406 0 0 0 4.295 12.193h30.408a19.453 19.453 0 0 0 0-24.386H8.796A19.406 19.406 0 0 0 4.501 24Z"/>`),
		g.Group(children),
	)
}