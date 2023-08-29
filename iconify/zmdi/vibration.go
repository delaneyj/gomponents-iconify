package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vibration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 256V128h43v128H0zm64 43V85h43v214H64zm405-171h43v128h-43V128zm-64 171V85h43v214h-43zM352 0q13 0 22.5 9.5T384 32v320q0 13-9.5 22.5T352 384H160q-13 0-22.5-9.5T128 352V32q0-13 9.5-22.5T160 0h192zm-11 341V43H171v298h170z"/>`),
		g.Group(children),
	)
}