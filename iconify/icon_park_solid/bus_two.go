package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBusTwo0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" d="M4 6a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v32a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6Z" clip-rule="evenodd"/><path stroke="#fff" d="M16 40H8v4h8v-4Zm24 0h-8v4h8v-4Z"/><path stroke="#000" d="M21 16h6M10 34h2m7 0h10M4 25h40M4 10h40m-8 24h2"/><path stroke="#fff" d="M4 6v8m40-8v8M4 21v8m40-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBusTwo0)"/>`),
		g.Group(children),
	)
}