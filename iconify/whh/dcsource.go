package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dcsource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 192h-96v96q0 13-9.5 22.5T864 320t-22.5-9.5T832 288v-96h-96q-13 0-22.5-9.5T704 160t9.5-22.5T736 128h96V32q0-13 9.5-22.5T864 0t22.5 9.5T896 32v96h96q13 0 22.5 9.5t9.5 22.5t-9.5 22.5T992 192zm32 256q0 27-18.5 45.5T960 512H640v256q0 27-19 45.5T575.5 832t-45-18.5T512 768V128q0-27 18.5-45.5t45-18.5T621 82.5t19 45.5v256h320q27 0 45.5 18.5T1024 448zM319.5 640q-26.5 0-45-18.5T256 576v-64H64q-27 0-45.5-18.5T0 448.5T18.5 403T64 384h192v-64q0-27 18.5-45.5t45-18.5t45.5 19t19 45v256q0 27-19 45.5T319.5 640z"/>`),
		g.Group(children),
	)
}