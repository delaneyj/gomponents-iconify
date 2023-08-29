package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Verticalborder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 960v-64h64v64h-64zm0-192h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zm-64-64h64v64h-64v-64zm128-64v64h-64v-64h64zm-64-128h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64V0zM768 0h64v64h-64V0zm0 512h-64v-64h64v64zM640 0h64v64h-64V0zm0 512h-64v-64h64v64zm-128-64v448h64v64H384v-64h64V64h-64V0h192v64h-64v384zm-128 64h-64v-64h64v64zM256 0h64v64h-64V0zm0 512h-64v-64h64v64zM128 0h64v64h-64V0zm0 512H64v-64h64v64zM0 576v-64h64v64H0zm0-192h64v64H0v-64zm0-128h64v64H0v-64zm0-128h64v64H0v-64zM0 0h64v64H0V0zm64 704H0v-64h64v64zm0 128H0v-64h64v64zm0 128H0v-64h64v64zm128-64v64h-64v-64h64zm128 64h-64v-64h64v64zm384 0h-64v-64h64v64zm128 0h-64v-64h64v64z"/>`),
		g.Group(children),
	)
}