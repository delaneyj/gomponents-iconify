package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 128h-42V0H85v128H43q-18 0-30.5 12.5T0 171v213h85v128h342V384h85V171q0-18-12.5-30.5T469 128zM128 43h256v85H128V43zm256 426H128V299h256v170zm85-128h-42v-85H85v85H43V171h426v170zm-42-128q0 9-6.5 15.5T405 235q-8 0-14.5-6.5T384 213q0-8 6.5-14.5T405 192q9 0 15.5 6.5T427 213zM149 320h214v43H149v-43zm0 64h214v43H149v-43z"/>`),
		g.Group(children),
	)
}