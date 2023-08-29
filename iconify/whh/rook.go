package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 1024H64q-26 0-45-18.5T0 960v-64q0-26 19-45t45-19h6q76-4 129-75t57-181q-27-8-45.5-28T192 512q0-14 20-35t44-29q-15-74-53.5-133T128 256q-26 0-45-18.5T64 192V64q0-26 19-45t45.5-19t45 19T192 64v32q0 13 9.5 22.5T224 128t22.5-9.5T256 96V64q0-26 19-45t45-19h128q27 0 45.5 19T512 64v32q0 13 9.5 22.5T544 128t22.5-9.5T576 96V64q0-26 19-45t45.5-19t45 19T704 64v128q0 27-18.5 45.5T640 256q-36 0-74.5 59T512 448q24 8 44 29t20 35q0 16-18.5 36T512 576q4 110 57 181t129 75h6q27 0 45.5 19t18.5 45v64q0 27-18.5 45.5T704 1024z"/>`),
		g.Group(children),
	)
}