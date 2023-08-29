package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixellance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 128V64h-64V0h128v128h-64zm0 64h-64v-64h64v64zm-128-64V64h64v64h-64zm-128 0h128v64H704v-64zm0 128h64v64h64v64H704v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v128H0V832h128v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64V192h64v64zM128 896H64v64h64v-64zm768-704v128h-64V192h64z"/>`),
		g.Group(children),
	)
}