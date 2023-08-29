package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelbattleaxe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 576v-64h64v-64H640v-64h64v-64h64v-64h64v-64h64v128h64v-64h64v320H704zm256-192H832v128h128V384zm0-256v64h-64v-64h64zm-64-64V0h128v128h-64V64h-64zm-64 64V64h64v64h-64zm-64 64v64h-64v64h-64v64h-64V256h-64v64h-64V0h320v64h-64v64h128v64h-64zM640 64H512v128h128V64zM448 512v-64h64v-64h64v64h64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v128H0V832h128v-64h64v-64h64v-64h64v-64h64v-64h64zM128 896H64v64h64v-64z"/>`),
		g.Group(children),
	)
}