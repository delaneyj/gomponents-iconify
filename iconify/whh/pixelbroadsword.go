package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelbroadsword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 256V64H768V0h256v256h-64zM768 128h-64V64h64v64zm-64 64h-64v-64h64v64zm-64 0v64h-64v-64h64zM512 320v-64h64v64h-64zm-64 64v-64h64v64h-64zm-64 64v-64h64v64h-64zm-64 64v-64h64v64h-64zm64 128h-64v-64h64v64zm64 64h-64v-64h64v64zm128-64v64h-64v-64h64zm64-64v64h-64v-64h64zm64-64v64h-64v-64h64zm64-64v64h-64v-64h64zm64-64v64h-64v-64h64zM512 768h-64v-64h64v64zm64 128H384v-64h128v-64h64v128zM256 576v-64h64v64h-64zm-64 0v64h-64V448h128v64h-64v64zm0 128v64h-64v-64h64zm0 192h64v64h-64v64H64v-64H0V832h64v-64h64v64H64v64h64v64h64v-64zm64-64h64v64h-64v-64zm64-64h64v64h-64v-64zm-64-64h64v64h-64v-64zm-64-64h64v64h-64v-64zm576-384v-64h64v64h-64zm-64 64v-64h64v64h-64zm-64 64v-64h64v64h-64zm-64 64v-64h64v64h-64zm-64 64v-64h64v64h-64zm-64 64v-64h64v64h-64zm448-448v64h-64v-64h64zm0 256h-64v-64h64v64zm64-64h-64v-64h64v64z"/>`),
		g.Group(children),
	)
}