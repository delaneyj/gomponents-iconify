package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Innerborders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 960v-64h64v64h-64zm0-192h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128H512v384h64v64H384v-64h64V512H64v64H0V384h64v64h384V64h-64V0h192v64h-64v384h384v-64h64v192h-64v-64zm0-256h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64V0zM768 0h64v64h-64V0zM640 0h64v64h-64V0zM256 0h64v64h-64V0zM128 0h64v64h-64V0zM0 256h64v64H0v-64zm0-128h64v64H0v-64zM0 0h64v64H0V0zm64 704H0v-64h64v64zm0 128H0v-64h64v64zm0 128H0v-64h64v64zm128 0h-64v-64h64v64zm128 0h-64v-64h64v64zm384 0h-64v-64h64v64zm128 0h-64v-64h64v64z"/>`),
		g.Group(children),
	)
}