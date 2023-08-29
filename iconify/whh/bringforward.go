package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bringforward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.488 1024h-512q-27 0-45.5-18.5t-18.5-45.5V704h64v256h512V448h-256v-64h256q27 0 45.5 18.5t18.5 45.5v512q0 27-18.5 45.5t-45.5 18.5zm-384-384h-512q-27 0-45.5-18.5T.488 576V64q0-27 18.5-45.5T64.488 0h512q27 0 45.5 18.5t18.5 45.5v512q0 27-18.5 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}