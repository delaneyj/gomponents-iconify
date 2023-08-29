package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelpotion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 896v64h-64v64H128v-64H64v-63H0V577h64v-64h64v-64h64v-64h-64v-63h512v63h-64v64h64v64h64v64h64v320h-64v-1zM128 768h64V640h-64v128zm512-256H512V384H256v128H128v64h512v-64zm0 256h-64v64h-64v64h64v-64h64v-64zM64 192h64v128H64V192zm640 0v128h-64V192h64zm-128 64H192v-64h-64v-64h64V65h64v128h256V65h64v63h64v64h-64v64zM256 0h256v64H256V0z"/>`),
		g.Group(children),
	)
}