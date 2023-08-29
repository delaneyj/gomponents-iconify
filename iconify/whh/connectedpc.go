package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connectedpc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 960H704q0 27-18.5 45.5T640 1024H384q-26 0-45-19t-19-45H64q-26 0-45-19T0 896V768q0-27 19-45.5T64 704h256q0-27 19-45.5t45-18.5v-64H192q-26 0-45-19t-19-45V64q0-27 19-45.5T192 0h640q27 0 45.5 18.5T896 64v448q0 27-18.5 45.5T832 576H640v64q27 0 45.5 18.5T704 704h256q27 0 45.5 19t18.5 45v128q0 26-18.5 45T960 960zM320 768H96q-13 0-22.5 9.5T64 800t9.5 22.5T96 832h224v-64zm512-320V64H192v384h640zM448 576v64h64v-64h-64zm192 128H384v256h256V704zm288 64H704v64h224q13 0 22.5-9.5T960 800t-9.5-22.5T928 768z"/>`),
		g.Group(children),
	)
}