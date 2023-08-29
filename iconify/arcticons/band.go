package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Band(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.2 27.9V8.4a3.9 3.9 0 0 0-3.9-3.9h0a3.9 3.9 0 0 0-3.9 3.9v19.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.2 14.387A15.602 15.602 0 1 1 8.4 27.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.2 19.18a11.7 11.7 0 1 1-3.9 8.72V8.4"/><circle cx="24" cy="27.9" r="7.8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}