package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaShieldOffOutline0"><g id="evaShieldOffOutline1"><path id="evaShieldOffOutline2" fill="currentColor" d="M4.71 3.29a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm7.59 16.39l-.3.17l-.3-.17A13.15 13.15 0 0 1 5 8.23v-.14L5.16 8L3.73 6.56A2 2 0 0 0 3 8.09v.14a15.17 15.17 0 0 0 7.72 13.2l.3.17a2 2 0 0 0 2 0l.3-.17a15.22 15.22 0 0 0 3-2.27l-1.42-1.42a12.56 12.56 0 0 1-2.6 1.94ZM20 6.34L13 2.4a2 2 0 0 0-2 0L7.32 4.49L8.78 6L12 4.15l7 3.94v.14a13 13 0 0 1-1.63 6.31L18.84 16A15.08 15.08 0 0 0 21 8.23v-.14a2 2 0 0 0-1-1.75Z"/></g></g>`),
		g.Group(children),
	)
}