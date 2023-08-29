package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Joystickarcade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024H64q-26 0-45-19T0 960V832q0-26 19-45t45-19h256q0-26 18.5-45t45.5-19h64V372q-57-20-92.5-69T320 192q0-79 56-135.5T512 0t136 56t56 136q0 62-35.5 111T576 372v332h64q27 0 45.5 19t18.5 45h256q26 0 45 19t19 45v128q0 27-19 45.5t-45 18.5zM128 640h64q26 0 45 19t19 45H64q0-26 18.5-45t45.5-19z"/>`),
		g.Group(children),
	)
}