package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5h-18m9-9v18m6-6l3-3l-3-3m-12 6l-3-3l3-3m3-3l3-3l3 3m-6 12l3 3l3-3"/>`),
		g.Group(children),
	)
}