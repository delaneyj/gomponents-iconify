package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardMembership(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v235q0 18-12.5 30.5T384 323h-85v106l-86-42l-85 42V323H43q-18 0-30.5-12.5T0 280V45q0-17 12.5-29.5T43 3h341zm0 277v-43H43v43h341zm0-107V45H43v128h341z"/>`),
		g.Group(children),
	)
}