package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelshield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 576V192H640v-64h256v448h-64zm-64-320v320h-64v192h-64v64H512v64H384v-64H256v-64h-64V576h-64V256h128v-65h384v65h128zM64 576H0V128h256v64H64v384zm64 192H64V576h64v192zm704-192v192h-64V576h64zm-64 256h-64v-64h64v64zm-64 64h-64v-64h64v64zm-64 64H512v-64h128v64zm-384 0v-64h128v64H256zm-64-64v-64h64v64h-64zm0-64h-64v-64h64v64zM512 64h128v64H512V64zm-128 64H256V64h128v64zm0-128h128v64H384V0zm128 1024H384v-64h128v64z"/>`),
		g.Group(children),
	)
}