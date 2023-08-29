package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Topborder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 960v-64h64v64h-64zm0-192h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zm-64-64h64v64h-64v-64zm128-64v64h-64v-64h64zm-64-128h64v64h-64v-64zm0-128h64v64h-64v-64zM512 64v64h-64V64H0V0h960v64H512zM64 192H0v-64h64v64zm0 128H0v-64h64v64zm0 128H0v-64h64v64zm64 64H64v-64h64v64zM0 576v-64h64v64H0zm64 128H0v-64h64v64zm0 128H0v-64h64v64zm0 128H0v-64h64v64zm128 0h-64v-64h64v64zm64-448h-64v-64h64v64zm64-64h64v64h-64v-64zm0 512h-64v-64h64v64zm192-192h-64v-64h64v64zm0-128h-64v-64h64v64zm0-128h-64v-64h64v64zm-64-320h64v64h-64v-64zm64 192h-64v-64h64v64zm64 64h64v64h-64v-64zm-64 384v64h-64v-64h64zm64 128h-64v-64h64v64zm-128 0h-64v-64h64v64zm256-512h64v64h-64v-64zm0 512h-64v-64h64v64zm128 0h-64v-64h64v64z"/>`),
		g.Group(children),
	)
}