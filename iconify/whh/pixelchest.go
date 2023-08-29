package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelchest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 832V448H768v384H192V448H64v384H0V192h64v192h832V192h64v640h-64zM512 448v128h-64V448h-64v192h192V448h-64zM832 64h64v128h-64V64zM704 192h64v128H192V192h64V64H128V0h704v64H704v128zm-640 0V64h64v128H64zm832 704H768v-64h128v64zm-704 0H64v-64h128v64z"/>`),
		g.Group(children),
	)
}