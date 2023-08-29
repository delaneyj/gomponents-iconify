package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1043 437q0-100-65-162t-171-62H487v448h320q106 0 171-62t65-162zm237 0q0 193-126.5 315T827 874H487v118h505q14 0 23 9t9 23v128q0 14-9 23t-23 9H487v192q0 14-9.5 23t-22.5 9H288q-14 0-23-9t-9-23v-192H32q-14 0-23-9t-9-23v-128q0-14 9-23t23-9h224V874H32q-14 0-23-9t-9-23V693q0-13 9-22.5t23-9.5h224V32q0-14 9-23t23-9h539q200 0 326.5 122T1280 437z"/>`),
		g.Group(children),
	)
}