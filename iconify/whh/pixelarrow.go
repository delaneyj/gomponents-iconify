package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelarrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 192V64H832V0h192v192h-64zm-256-64V64h128v64H704zm64 64v64h64v64h64v64h-64v-64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v64h-64v128h-64v-64h-64v-64H64v-64h128v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64h64v-64h-64v-64h64v64h64zM128 960H64v-64h64v64zm832-768v128h-64V192h64zm-768 832h-64v-64h64v64zm-192 0v-64h64v64H0zm0-128v-64h64v64H0z"/>`),
		g.Group(children),
	)
}