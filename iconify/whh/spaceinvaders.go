package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spaceinvaders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 704V512H832v192h-64v-64H256v64h-64V512H64v192H0V448h64V320h128V192h128v-64h64v64h256v-64h64v64h128v128h128v128h64v256h-64zM384 320H256v128h128V320zm384 0H640v128h128V320zM576 768v-64h192v64H576zm-128 0H256v-64h192v64zM256 64h64v64h-64V64zM128 0h128v64H128V0zm640 64v64h-64V64h64zM896 0v64H768V0h128z"/>`),
		g.Group(children),
	)
}