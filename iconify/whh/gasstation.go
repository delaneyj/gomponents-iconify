package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gasstation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.405 127h-64q-26 0-45 19t-19 45v448q0 53-37.5 90.5t-90.5 37.5h-64v128q27 0 45.5 19t18.5 45.5t-18.5 45t-45.5 18.5h-576q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19V64q0-27 18.5-45.5t45.5-18.5h448q26 0 45 18.5t19 45.5v575q27 0 45.5-18.5t18.5-45.5V127q0-53 37.5-90t90.5-37h128q27 0 45.5 18.5t18.5 45t-18.5 45t-45.5 18.5zm-384 64q0-26-18.5-45t-45.5-19h-320q-26 0-45 19t-19 45v128q0 27 18.5 45.5t45.5 18.5h320q27 0 45.5-18.5t18.5-45.5V191z"/>`),
		g.Group(children),
	)
}