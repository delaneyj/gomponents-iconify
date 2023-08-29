package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 704H704q0 27-18.5 45.5T640 768H384q-27 0-45.5-18.5T320 704H64q-27 0-45.5-19T0 640V512q0-27 18.5-45.5T64 448h256q0-27 18.5-45.5T384 384V64q0-27 18.5-45.5T448 0h128q27 0 45.5 18.5T640 64v320q27 0 45.5 18.5T704 448h256q27 0 45.5 18.5T1024 512v128q0 26-18.5 45T960 704zM320 512H96q-13 0-22.5 9.5T64 544t9.5 22.5T96 576h224v-64zM512 96q0-13-9.5-22.5T480 64t-22.5 9.5T448 96v288h64V96zm128 352H384v256h256V448zm288 64H704v64h224q13 0 22.5-9.5T960 544t-9.5-22.5T928 512z"/>`),
		g.Group(children),
	)
}