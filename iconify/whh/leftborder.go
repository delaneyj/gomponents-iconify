package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leftborder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 960v-64h64v64h-64zm0-192h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zm-64-64h64v64h-64v-64zm128-64v64h-64v-64h64zm-64-128h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64V0zM768 0h64v64h-64V0zm0 512h-64v-64h64v64zM640 0h64v64h-64V0zm0 512h-64v-64h64v64zM448 128V64h64v64h-64zM384 0h64v64h-64V0zm128 0h64v64h-64V0zM384 512h-64v-64h64v64zM256 0h64v64h-64V0zm0 512h-64v-64h64v64zM128 0h64v64h-64V0zm0 512H64v448H0V0h64v448h64v64zm64 448h-64v-64h64v64zm128 0h-64v-64h64v64zm192-192h-64v-64h64v64zm0-128h-64v-64h64v64zm0-128h-64v-64h64v64zm-64-320h64v64h-64v-64zm64 192h-64v-64h64v64zm0 448v64h-64v-64h64zm64 128h-64v-64h64v64zm-128 0h-64v-64h64v64zm256 0h-64v-64h64v64zm128 0h-64v-64h64v64z"/>`),
		g.Group(children),
	)
}