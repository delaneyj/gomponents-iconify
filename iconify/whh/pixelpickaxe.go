package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelpickaxe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 640V384h64v256h-64zm-64-320h64v64h-64v-64zm-64-64h-64v-64h-64v-64h128V64h128v128h-64v128h-64v-64zM640 64h64v64h-64V64zm-64 128H384v-64h192v64zm64 64h-64v-64h64v64zm64 64h64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v128H0V832h128v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v64zM128 896H64v64h64v-64zm704-448h-64v-64h64v64zm64 192h-64V448h64v192zm64 64h-64v-64h64v64zM320 128V64h64v64h-64zM384 0h256v64H384V0z"/>`),
		g.Group(children),
	)
}