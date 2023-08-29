package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21q-1.875 0-3.513-.7t-2.862-1.925Q3.4 17.15 2.7 15.512T2 12q0-1.875.7-3.513t1.925-2.862Q5.85 4.4 7.488 3.7T11 3q.525 0 1.012.063T13 3.25V5.3q-.5-.15-.988-.225T11 5Q8.05 5 6.025 7.025T4 12q0 2.95 2.025 4.975T11 19q2.95 0 4.975-2.025T18 12q0-.275-.025-.5T17.9 11h2.05q.05.275.05.5v.5q0 1.875-.7 3.513t-1.925 2.862Q16.15 19.6 14.512 20.3T11 21Zm2.8-4.8L10 12.4V7h2v4.6l3.2 3.2l-1.4 1.4ZM18 9V6h-3V4h3V1h2v3h3v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}