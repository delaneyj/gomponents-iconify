package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bottomborder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M0 960v-64h448v-64h64v64h448v64H0zm896-192h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zm-64-64h64v64h-64v-64zm128-64v64h-64v-64h64zm-64-128h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64V0zM768 512h-64v-64h64v64zm0-512h64v64h-64V0zM640 512h-64v-64h64v64zm0-512h64v64h-64V0zM448 192h64v64h-64v-64zm0 128h64v64h-64v-64zm0 128h64v64h-64v-64zm64 320h-64v-64h64v64zm-64-192h64v64h-64v-64zm-64-64h-64v-64h64v64zm64-384V64h64v64h-64zM384 0h64v64h-64V0zm128 0h64v64h-64V0zM256 0h64v64h-64V0zm0 512h-64v-64h64v64zM128 0h64v64h-64V0zm0 512H64v-64h64v64zM0 576v-64h64v64H0zm0-192h64v64H0v-64zm0-128h64v64H0v-64zm0-128h64v64H0v-64zM0 0h64v64H0V0zm64 704H0v-64h64v64zm0 128H0v-64h64v64z"/>`),
		g.Group(children),
	)
}