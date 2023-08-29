package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Router(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M367 62q-40-36-90-36t-89 36l-17-17Q215 0 277 0t107 45zm-19 17l-17 17q-22-21-54-21t-53 21l-17-17q30-30 70.5-30T348 79zm-7 134q18 0 30.5 12.5T384 256v85q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341v-85q0-18 12.5-30.5T43 213h213v-85h43v85h42zM107 320v-43H64v43h43zm74 0v-43h-42v43h42zm75 0v-43h-43v43h43z"/>`),
		g.Group(children),
	)
}