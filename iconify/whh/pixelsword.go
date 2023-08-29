package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelsword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 192v64h-64v-64h64V64H832v64h-64V64h64V0h192v192h-64zm-192 0h-64v-64h64v64zm-64 0v64h-64v-64h64zM576 320v-64h64v64h-64zm-64 64v-64h64v64h-64zm-64 64v-64h64v64h-64zm-64 64v-64h64v64h-64zm-64 128V512h64v128h-64zm0 64h64v64h-64v64h-64v64h-64v128H0V832h128v-64h64v-64h64v-64h64v64zM128 896H64v64h64v-64zm256-256h128v64H384v-64zm192-64v64h-64v-64h64zm64-64v64h-64v-64h64zm64-64v64h-64v-64h64zm64-64v64h-64v-64h64zm64-64v64h-64v-64h64zM576 704v64h-64v-64h64zm64 192H512v-64h64v-64h64v128zm-128-64H384v-64h128v64zM192 640V512h64v128h-64zm64-192h64v64h-64v-64zm-64 64h-64V384h128v64h-64v64zm704-192h-64v-64h64v64z"/>`),
		g.Group(children),
	)
}