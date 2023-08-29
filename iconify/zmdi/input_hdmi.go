package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InputHdmi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M277 109h22v128l-64 128v64H64v-64L0 237V109h21V45q0-17 12.5-29.5T64 3h171q17 0 29.5 12.5T277 45v64zM64 45v64h43V67h21v42h43V67h21v42h43V45H64z"/>`),
		g.Group(children),
	)
}