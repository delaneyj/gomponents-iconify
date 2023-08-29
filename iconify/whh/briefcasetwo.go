package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcasetwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024H64q-26 0-45-18.5T0 960V448h192v64h-64q0 27 18.5 45.5T192 576h128q27 0 45.5-18.5T384 512h-64v-64h384v64h-64q0 27 19 45.5t45 18.5h128q27 0 45.5-18.5T896 512h-64v-64h192v512q0 27-18.5 45.5T960 1024zM832 320H704q-26 0-45 19t-19 45H384q0-26-19-45t-45-19H192q-27 0-45.5 19T128 384H0V256q0-26 19-45t45-19h192v-64q0-53 37.5-90.5T384 0h256q53 0 90.5 37.5T768 128v64h192q26 0 45 19t19 45v128H896q0-26-18.5-45T832 320zM576 128H448q-26 0-45 18.5T384 192h256q0-26-18.5-45T576 128z"/>`),
		g.Group(children),
	)
}