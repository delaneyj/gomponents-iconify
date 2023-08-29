package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Harddrivealt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024H128q-53 0-90.5-37.5T0 896V736q0-33 28.5-210t68-351.5T160 0h704q24 0 63.5 174.5t68 351.5t28.5 210v160q0 53-37.5 90.5T896 1024zm64-256q0-26-19-45t-45-19H128q-27 0-45.5 18.5T64 768v128q0 27 19 45.5t45 18.5h768q27 0 45.5-18.5T960 896V768zM192 896q-27 0-45.5-19T128 831.5t18.5-45T192 768t45.5 18.5t18.5 45t-18.5 45.5t-45.5 19z"/>`),
		g.Group(children),
	)
}