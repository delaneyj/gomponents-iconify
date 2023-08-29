package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardTravel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 88q18 0 30.5 12.5T427 131v234q0 18-12.5 30.5T384 408H43q-18 0-30.5-12.5T0 365V131q0-18 12.5-30.5T43 88h64V45q0-17 12.5-29.5T149 3h128q18 0 30.5 12.5T320 45v43h64zM149 45v43h128V45H149zm235 320v-42H43v42h341zm0-106V131h-64v42h-43v-42H149v42h-42v-42H43v128h341z"/>`),
		g.Group(children),
	)
}