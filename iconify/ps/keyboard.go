package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 64H43q-18 0-30.5 12.5T0 107v213q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30V107q0-18-12.5-30.5T469 64zM43 320V107h426v213H43zm85-64h256v43H128v-43zm277-64h-42v43h85V128h-43v64zm0 64h43v43h-43v-43zm-341 0h43v43H64v-43zm0-64h85v43H64v-43zm235 0h42v43h-42v-43zm-64 0h42v43h-42v-43zm-64 0h42v43h-42v-43zm149-64h43v43h-43v-43zm-64 0h43v43h-43v-43zm-64 0h43v43h-43v-43zm-64 0h43v43h-43v-43zm-64 0h43v43H64v-43z"/>`),
		g.Group(children),
	)
}