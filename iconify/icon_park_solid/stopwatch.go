package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSStopwatch0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44c9.389 0 17-7.611 17-17s-7.611-17-17-17S7 17.611 7 27s7.611 17 17 17Z"/><path stroke="#fff" stroke-linecap="round" d="M31 4H17m21 6l-3 3"/><path stroke="#000" stroke-linecap="round" d="M24 18v9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSStopwatch0)"/>`),
		g.Group(children),
	)
}