package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeSearchedFor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.55 20.45L14.8 14.7q-.8.65-1.775.975T11 16q-1.05 0-2-.313T7.25 14.8l1.45-1.45q.5.3 1.075.475T11 14q1.875 0 3.187-1.312T15.5 9.5q0-1.875-1.313-3.188T11 5Q9.275 5 8.012 6.163T6.55 9.05L7.8 7.8l1.4 1.4l-3.7 3.7l-3.7-3.7l1.4-1.4l1.35 1.3q.15-2.575 2-4.337T11 3q2.725 0 4.612 1.888T17.5 9.5q0 1.05-.325 2.05T16.2 13.3l5.75 5.75l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}