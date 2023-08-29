package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M864 1024H544q-13 0-22.5-9.5T512 992v-32H384v32q0 13-9.5 22.5T352 1024H32q-13 0-22.5-9.5T0 992L192 32q0-13 9.5-22.5T224 0h128q13 0 22.5 9.5T384 32v32h128V32q0-13 9.5-22.5T544 0h128q13 0 22.5 9.5T704 32l192 960q0 13-9.5 22.5T864 1024zM512 192H384v256h128V192zM384 576v256h128V576H384z"/>`),
		g.Group(children),
	)
}