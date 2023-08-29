package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024H64q-26 0-45-19T0 959.5t19-45T64 896h896q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19zM640 832H384q-106 0-181-75t-75-181V64q0-26 19-45t45.5-19t45 19T256 64v512q0 53 37.5 90.5T384 704h256q53 0 90.5-37.5T768 576V64q0-26 19-45t45.5-19t45 18.5T896 64v512q0 106-75 181t-181 75z"/>`),
		g.Group(children),
	)
}