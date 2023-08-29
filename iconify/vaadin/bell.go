package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 14h4s.1 2-2 2s-2-2-2-2zm6.7-2.6c-.5-.2-.7-.7-.7-1.2V5s-.2-2.4-3-2.9V1s.1-1-1-1s-1 1-1 1v1.1C4.2 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v1h12v-1l-1.3-.6zM6 4.8V12H4c.8 0 1-1 1-1V5s0-.8.7-1.4C6.4 2.9 7 3 7 3s-1 .7-1 1.8z"/>`),
		g.Group(children),
	)
}