package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emergency(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm0-576q0-27-18.5-45.5t-45.5-18.5h-192V192q0-27-18.5-45.5t-45.5-18.5h-128q-27 0-45.5 18.5t-18.5 45.5v192h-192q-26 0-45 18.5t-19 45.5v128q0 27 19 45.5t45 18.5h192v192q0 27 19 45.5t45 18.5h128q27 0 45.5-18.5t18.5-45.5V640h192q27 0 45.5-18.5t18.5-45.5V448z"/>`),
		g.Group(children),
	)
}