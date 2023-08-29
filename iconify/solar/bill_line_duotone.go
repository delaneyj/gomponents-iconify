package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BillLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M2 2h20"/><path d="M4 2h16v10.277c0 1.34 0 2.01-.268 2.601c-.268.591-.773 1.032-1.781 1.915l-2 1.75c-1.883 1.647-2.824 2.47-3.951 2.47c-1.127 0-2.068-.823-3.951-2.47l-2-1.75c-1.009-.883-1.513-1.324-1.78-1.915C4 14.288 4 13.618 4 12.278V2Z" opacity=".5"/><path stroke-linecap="round" d="M8.5 13h7m-7-5h7"/></g>`),
		g.Group(children),
	)
}