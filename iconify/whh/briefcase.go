package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024H64q-26 0-45-19T0 960V576h384v64q0 26 19 45t45 19h128q27 0 45.5-19t18.5-45v-64h384v384q0 27-19 45.5t-45 18.5zM576 575.5q0 26.5-18.5 45.5t-45 19t-45.5-19t-19-45.5t19-45t45.5-18.5t45 18.5t18.5 45zm0-127.5H448q-26 0-45 18.5T384 512H0V320q0-27 18.5-45.5T64 256h192v-64q0-80 56-136T448 0h128q80 0 136 56t56 136v64h192q26 0 45 18.5t19 45.5v192H640q0-27-18.5-45.5T576 448zm64-256q0-27-18.5-45.5T576 128H448q-26 0-45 18.5T384 192v64h256v-64z"/>`),
		g.Group(children),
	)
}