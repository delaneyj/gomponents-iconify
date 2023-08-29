package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M416 224H208v120.619h-22.154v-41A87.716 87.716 0 0 0 98.229 216H48v-64H16v344h32v-47.743l416 3.328V496h32V304a80.091 80.091 0 0 0-80-80ZM48 248h50.229a55.68 55.68 0 0 1 55.617 55.617v41H48Zm416 171.584l-416-3.328v-39.637h416Zm0-74.965H240V256h176a48.055 48.055 0 0 1 48 48Z"/>`),
		g.Group(children),
	)
}