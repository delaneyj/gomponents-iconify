package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M290 347H176q-24 0-40-16q-17-19-17-41v-40h162q22 0 38-16t16-38t-16-37q-15-16-38-16H119V59q0-24-16-40Q86 2 62 2Q37 2 21 18Q4 35 4 59v231q0 71 51 122q49 50 121 50h114q24 0 41-17t17-40q0-24-17-41q-19-17-41-17z"/>`),
		g.Group(children),
	)
}