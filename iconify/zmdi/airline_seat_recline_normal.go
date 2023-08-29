package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatReclineNormal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M76.5 75.5Q64 63 64 45.5t12.5-30t30-12.5t30 12.5t12.5 30t-12.5 30t-30 12.5t-30-12.5zM43 301q0 27 18.5 45.5T107 365h128v43H107q-44 0-75.5-31.5T0 301V109h43v192zm298 87l-30 31l-75-75H128q-27 0-45.5-18.5T64 280V157q0-20 14-34t34-14h1q10 0 20 5q9 4 15 11l30 33q17 19 45 31.5t54 11.5v47q-29 0-61-13.5T160 201v79h73z"/>`),
		g.Group(children),
	)
}