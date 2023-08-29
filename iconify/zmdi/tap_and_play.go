package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapAndPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 325q62 0 105.5 44T149 475h-42q0-44-31.5-75.5T0 368v-43zm0 86q27 0 45.5 18.5T64 475H0v-64zm0-171q97 0 166 68.5T235 475h-43q0-80-56-136T0 283v-43zM320 6q18 0 30.5 12T363 48v363q0 17-12.5 29.5T320 453h-44q-4-45-21-85h65V91H107v128q-20-8-43-14V48q0-18 12.5-30.5T107 5z"/>`),
		g.Group(children),
	)
}