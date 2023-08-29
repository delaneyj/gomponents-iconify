package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 640q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h256q26 0 45 19t19 45v192q0 106-75 181t-181 75H256q-106 0-181-75T0 768V256Q0 150 75 75T256 0h256q106 0 181 75t75 181q0 27-19 45.5T703.5 320t-45-18.5T640 256q0-53-37.5-90.5T512 128H256q-53 0-90.5 37.5T128 256v512q0 53 37.5 90.5T256 896h256q53 0 90.5-37.5T640 768V640H448z"/>`),
		g.Group(children),
	)
}