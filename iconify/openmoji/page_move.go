package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageMove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="m36 15.916l16 16.117v23.968H20V15.916h16"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m36 15.916l16 16.117v23.968H20V15.916h16"/><path d="m36 15.916l-.034 16.117h10.573m12.947-.986l5.216 5.216m-5.216 5.258l5.216-5.216m-52 5.216l-5.215-5.216m5.215-5.258l-5.215 5.216"/></g>`),
		g.Group(children),
	)
}