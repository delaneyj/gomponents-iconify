package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M41 192q0 27 19 46.5t47 19.5h85v41h-85q-44 0-75.5-31.5T0 192t31.5-75.5T107 85h85v41h-85q-28 0-47 19.5T41 192zm87 21v-42h171v42H128zM320 85q44 0 75.5 31.5T427 192t-31.5 75.5T320 299h-85v-41h85q27 0 46.5-19.5T386 192t-19.5-46.5T320 126h-85V85h85z"/>`),
		g.Group(children),
	)
}