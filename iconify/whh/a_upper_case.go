package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M703.5 1024q-26.5 0-45-18.5T640 960V672q0-13-9.5-22.5T608 640H160q-13 0-22.5 9.5T128 672v288q0 27-19 45.5T64 1024t-45-18.5T0 960V256Q0 150 75 75T256 0h256q106 0 181 75t75 181v704q0 27-19 45.5t-45.5 18.5zM640 256q0-53-37.5-90.5T512 128H256q-53 0-90.5 37.5T128 256v224q0 13 9.5 22.5T160 512h448q13 0 22.5-9.5T640 480V256z"/>`),
		g.Group(children),
	)
}