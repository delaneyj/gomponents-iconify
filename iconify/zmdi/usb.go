package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M218 133h85v86h-21v42q0 18-12.5 30.5T239 304h-64v65q26 13 26 42q0 19-14 33t-33.5 14t-33-14t-13.5-33q0-29 25-42v-65H68q-17 0-29.5-12.5T26 261v-44Q0 204 0 176q0-19 14-33t33-14t33 14t14 33q0 28-26 41v44h64V91H90l64-86l64 86h-43v170h64v-42h-21v-86z"/>`),
		g.Group(children),
	)
}