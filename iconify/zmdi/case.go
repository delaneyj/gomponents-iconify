package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Case(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 88q18 0 30.5 12.5T427 131v234q0 18-12.5 30.5T384 408H43q-18 0-30.5-12.5T0 365V131q0-18 12.5-30.5T43 88h85V45q0-17 12.5-29.5T171 3h85q18 0 30.5 12.5T299 45v43h85zm-128 0V45h-85v43h85z"/>`),
		g.Group(children),
	)
}