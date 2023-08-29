package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 320v-64h341v128h-64v-64H107v64H43v-64zm320-171h64v64h-64v-64zM0 149h64v64H0v-64zm320 64H107V43q0-18 12.5-30.5T149 0h128q18 0 30.5 12.5T320 43v170z"/>`),
		g.Group(children),
	)
}