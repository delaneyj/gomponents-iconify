package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelpotionalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 896V576h64v320h-64zm-64-384H512V384H256v128H128v64H64v-64h64v-64h64v-64h-64v-64h512v64h-64v64h64v64h64v64h-64v-64zM0 896V576h64v320H0zm64-704h64v128H64V192zm640 0v128h-64V192h64zm-128 64H192v-64h-64v-64h64V64h64v128h256V64h64v64h64v64h-64v64zM256 0h256v64H256V0zm-64 960h384v-64h128v64h-64v64H128v-64H64v-64h128v64z"/>`),
		g.Group(children),
	)
}