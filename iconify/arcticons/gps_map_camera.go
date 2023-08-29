package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpsMapCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5c-8.185 0-14.82 6.635-14.82 14.82h0v1.08c.6 8.12 7.34 14.65 14.82 23.1c7.81-8.82 14.82-15.5 14.82-24.18h0C38.82 11.135 32.185 4.5 24 4.5h0Zm10 14.83c0 5.523-4.477 10-10 10s-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10Zm-4.05-3.906a2.044 2.044 0 1 1-4.089 0a2.044 2.044 0 0 1 4.09 0Z"/>`),
		g.Group(children),
	)
}