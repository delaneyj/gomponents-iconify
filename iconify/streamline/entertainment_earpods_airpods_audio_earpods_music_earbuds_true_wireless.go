package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentEarpodsAirpodsAudioEarpodsMusicEarbudsTrueWireless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5a3 3 0 0 0-2 .76v4.48a3 3 0 0 0 2 .76a2.74 2.74 0 0 0 .5 0v5.79a1.25 1.25 0 0 0 2.5 0V3.5a3 3 0 0 0-3-3Zm9 0a3 3 0 0 1 2 .76v4.48a3 3 0 0 1-2 .76a2.74 2.74 0 0 1-.5 0v5.79a1.25 1.25 0 0 1-2.5 0V3.5a3 3 0 0 1 3-3Z"/>`),
		g.Group(children),
	)
}