package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelelixir(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 832v64h-64v64h-64v64H192v-64h-64v-64H64v-64h64v64h64v-64h64v64h64v64h64v-64h-64V576h64V384h-64v-64h64v-64h-64v-64h64V64h-64V0h128v64h64v128h-64v64h64v64h-64v64h64v64h64v64h64v64h64v64h64v192h-64zm-64-128h-64v128h-64v64h64v-64h64V704zM320 576h-64v-64h-64v-64h64v-64h64v192zm-192 0v64H64v-64h64v-64h64v64h-64zm64 256h-64V640h64v192zm0-192v-64h64v64h-64zM0 832V640h64v192H0zm256-512v-64h64v64h-64zm-64-128h64v64h-64v-64zm64-128h64v128h-64V64zm320 128v64h-64v-64h64z"/>`),
		g.Group(children),
	)
}