package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Downright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.63 193h-128q-26 0-45 18.5t-19 45.5v324l-567-562q-18-19-45-19t-46 19l-91 90q-19 19-19 45.5t19 45.5l572 568h-335q-26 0-45 18.5t-19 44.5v128q0 27 19 45.5t45 18.5h704q27 0 45.5-18.5t18.5-45.5V256q0-26-18.5-44.5t-45.5-18.5z"/>`),
		g.Group(children),
	)
}