package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Regex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.88 8.4H6.5a2 2 0 0 0-2 2v27.2a2 2 0 0 0 2 2h9.38m16.24 0h9.38a2 2 0 0 0 2-2V10.4a2 2 0 0 0-2-2h-9.38m-3.78 22.05V14.79m-6.78 11.75l13.56-7.83m0 7.83l-13.56-7.83"/><circle cx="15.33" cy="30.45" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}