package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elevator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSElevator0"><g fill="none"><path fill="#fff" d="M42 41V7H6v34h36Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 7v34m0-34H4h2v34M42 7h2m-2 34h2m-2 0H6m0 0H4M24 7v34"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 20v8m-3-5l3-3l3 3m-21 5v-8m-3 5l3 3l3-3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSElevator0)"/>`),
		g.Group(children),
	)
}