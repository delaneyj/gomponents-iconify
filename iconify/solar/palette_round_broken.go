package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteRoundBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2 8V6a4 4 0 1 1 8 0v12a4 4 0 0 1-8 0v-6"/><path stroke="currentColor" stroke-width="1.5" d="m10 8.243l3.314-3.314a4 4 0 1 1 5.657 5.657L9.306 20.25"/><path fill="currentColor" d="M18 22v-.75v.75Zm0-8v.75V14Zm4 4h-.75h.75Zm-9 4.75a.75.75 0 0 0 0-1.5v1.5Zm4-1.5a.75.75 0 0 0 0 1.5v-1.5Zm-1.5-6.5H18v-1.5h-2.5v1.5ZM21.25 18A3.25 3.25 0 0 1 18 21.25v1.5A4.75 4.75 0 0 0 22.75 18h-1.5Zm1.5 0A4.75 4.75 0 0 0 18 13.25v1.5A3.25 3.25 0 0 1 21.25 18h1.5ZM13 21.25H6v1.5h7v-1.5Zm5 0h-1v1.5h1v-1.5Z"/><path stroke="currentColor" stroke-width="1.5" d="M7 18a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		g.Group(children),
	)
}