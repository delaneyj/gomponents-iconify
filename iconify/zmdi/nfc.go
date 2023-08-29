package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nfc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3h341zm0 384V45H43v342h341zM341 88v256H85V88h86v43h-43v170h171V131h-64v48q21 12 21 37q0 18-12.5 30.5t-30 12.5t-30-12.5T171 216q0-24 21-37v-48q0-18 12.5-30.5T235 88h106z"/>`),
		g.Group(children),
	)
}