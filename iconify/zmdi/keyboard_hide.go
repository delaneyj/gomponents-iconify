package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardHide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 0q18 0 30.5 12.5T427 43v213q0 18-12.5 30.5T384 299H43q-18 0-30.5-12.5T0 256V43q0-18 12.5-30.5T43 0h341zM192 64v43h43V64h-43zm0 64v43h43v-43h-43zm-64-64v43h43V64h-43zm0 64v43h43v-43h-43zm-21 43v-43H64v43h43zm0-64V64H64v43h43zm192 149v-43H128v43h171zm0-85v-43h-43v43h43zm0-64V64h-43v43h43zm64 64v-43h-43v43h43zm0-64V64h-43v43h43zM213 427l-85-86h171z"/>`),
		g.Group(children),
	)
}