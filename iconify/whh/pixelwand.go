package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelwand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 384v64h-64v64h-64v-64h-64v-64h-64v-64h-64v-64h-64v-64h-64v-64h64V64h64V0h256v64h64v64h64v256h-64zM832 64H704v64h128V64zm64 256h-64v64h64v-64zm-384 0h-64v-64h64v64zm-64 192h-64V320h64v192zm0 128h-64v-64h64v64zm0-128h64v64h-64v-64zm192 64h64v64H512v-64h128zm0-128h64v64h-64v-64zm-64-64h64v64h-64v-64zm-64-64h64v64h-64v-64zm256 192v64h-64v-64h64zm192 0v64h-64v-64h64zm-64 128H768v-64h128v64zm-384 64h-64v-64h64v64zm-128 64v-64h64v64h-64zm-64 64v-64h64v64h-64zm0-256v-64h64v64h-64zm-64 64v-64h64v64h-64zm-64 64v-64h64v64h-64zm192-448V128h64v128h-64zm64-192h64v64h-64V64zM256 832h64v64h-64v-64zm0 64v64h-64v64H64v-64h128v-64h64zM128 704h64v64h-64v-64zM64 960H0V832h64v-64h64v64H64v128z"/>`),
		g.Group(children),
	)
}