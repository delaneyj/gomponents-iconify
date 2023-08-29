package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Point(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPoint0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 41C13.5 38.25 8.001 30 5 25c-3-5 3.313-9.688 7-6l4 4V7.5a3.5 3.5 0 1 1 7 0V16a3.5 3.5 0 1 1 7 0v1.5a3.5 3.5 0 0 1 7-.002V22.5a3.5 3.5 0 0 1 7 0v8.744c0 1.155-.262 2.3-.913 3.254C41.917 36.212 39.602 39.035 36 41c-5.5 3-11.5 2.75-17 0Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPoint0)"/>`),
		g.Group(children),
	)
}