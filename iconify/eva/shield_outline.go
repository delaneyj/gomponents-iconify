package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaShieldOutline0"><g id="evaShieldOutline1"><path id="evaShieldOutline2" fill="currentColor" d="M12 21.85a2 2 0 0 1-1-.25l-.3-.17A15.17 15.17 0 0 1 3 8.23v-.14a2 2 0 0 1 1-1.75l7-3.94a2 2 0 0 1 2 0l7 3.94a2 2 0 0 1 1 1.75v.14a15.17 15.17 0 0 1-7.72 13.2l-.3.17a2 2 0 0 1-.98.25Zm0-17.7L5 8.09v.14a13.15 13.15 0 0 0 6.7 11.45l.3.17l.3-.17A13.15 13.15 0 0 0 19 8.23v-.14Z"/></g></g>`),
		g.Group(children),
	)
}