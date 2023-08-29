package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelbastardsword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 192v64h-64v-64h64V64H832v64h-64V64h64V0h192v192h-64zm-192 0h-64v-64h64v64zm-64 0v64h-64v-64h64zM576 320v-64h64v64h-64zm-64 64v-64h64v64h-64zm-128 64v-64h128v64H384zm-128 64v64h-64V448h192v64H256zm-64 192v-64h64v-64h128v64h64v128h-64v64h-64v64h-64v64h-64v64H64v-64h128v-64h-64v-64H64v128H0V832h64v-64h64v-64h64zm320-64h64v192H448v-64h64V640zm128-128v128h-64V512h64zm64-64v64h-64v-64h64zm64-64v64h-64v-64h64zm64-64v64h-64v-64h64zm64 0h-64v-64h64v64z"/>`),
		g.Group(children),
	)
}