package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LookFourSat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.47 19.28A23.42 23.42 0 0 1 28.7 7.8m.46-.21l-23.43 11m34.74.69L31 40.85"/><circle cx="34.54" cy="13.69" r="8.18" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31 21.22L14 42.27M26.2 14l-10.36 7.07"/>`),
		g.Group(children),
	)
}