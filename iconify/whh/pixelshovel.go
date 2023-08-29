package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelshovel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 256V64h64v192h-64zM640 128V64h128v64H640zm-64 64v-64h64v64h-64zm64 192h64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v128H0V832h128v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v64zM128 896H64v64h64v-64zm768-512v64h-64v-64h64zm64-128v128h-64V256h64zM832 512H704v-64h128v64zM512 320V192h64v128h-64zM768 0h192v64H768V0z"/>`),
		g.Group(children),
	)
}