package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tabletscreensize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.226 768h-896q-26 0-45-18.5t-19-45.5V64q0-26 19-45t45-19h896q27 0 45.5 19t18.5 45v640q0 27-18.5 45.5t-45.5 18.5zm0-704h-896v640h896V64zm-704 128l576 320l64-64v192h-192l64-64l-576-320l-64 64V128h192z"/>`),
		g.Group(children),
	)
}