package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelaxe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 320v192h-64v64h-64v-64h64V384h-64v-64H768v64h-64v192h192v64H640V512h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v128H0V832h128v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64H384V128h64V64h64v64h-64v192h192v-64h64V128h128V64h128v128h-64v128h128zM128 896H64v64h64v-64zM640 64H512V0h192v128h-64V64z"/>`),
		g.Group(children),
	)
}