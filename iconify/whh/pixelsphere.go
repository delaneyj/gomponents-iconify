package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelsphere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 704v128h-64v64h-64v64H704v64H320v-64H192v-64h-64v-64H64V704H0V320h64V192h64v-64h64V64h128V0h384v64h128v64h64v64h64v128h64v384h-64zM576 256h-64v-64H320v64h-64v64h64v64h192v-64h64v-64zm192 128h-64v64h64v-64zm64 320h-64v64h-64v64h64v-64h64v-64z"/>`),
		g.Group(children),
	)
}