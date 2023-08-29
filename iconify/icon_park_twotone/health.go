package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Health(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHealth0"><g fill="#555" stroke="#fff" stroke-width="4"><path d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M27.3 12c-1.823 0-3.3 1.435-3.3 3.204c0 3.205 3.9 6.118 6 6.796c2.1-.678 6-3.59 6-6.796C36 13.434 34.523 12 32.7 12a3.326 3.326 0 0 0-2.7 1.362A3.326 3.326 0 0 0 27.3 12Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHealth0)"/>`),
		g.Group(children),
	)
}