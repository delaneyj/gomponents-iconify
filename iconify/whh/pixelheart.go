package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelheart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 448v128h-64v64h-64v-64h-64v64h-64v64h-64v64h-64v64h64v64h-64v64h-64v-64h-64v-64h-64v-64h-64v-64h-64v-64h-64v-64H64V448H0V128h64V64h64V0h192v64h64v64h64v64h64v-64h64V64h64V0h192v64h64v64h64v320h-64zM256 64h-64v64h-64v64H64v64h64v-64h64v-64h64V64zm512 0h-64v64h-64v64h-64v64h64v-64h64v-64h64V64zM640 768v64h-64v-64h64zm64-64v64h-64v-64h64zm64-64v64h-64v-64h64z"/>`),
		g.Group(children),
	)
}