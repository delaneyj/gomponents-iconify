package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataArrival(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDataArrival0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" d="M6 42h36V6H30c-1.324 3.159-3.324 4.738-6 4.738S19.324 9.158 18 6H6v36Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="m15 24l6 6l12-12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDataArrival0)"/>`),
		g.Group(children),
	)
}