package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 768H128q-53 0-90.5-37.5T0 640V128q0-53 37.5-90.5T128 0h768q53 0 90.5 37.5T1024 128v512q0 53-37.5 90.5T896 768zM448 160q0-13-9.5-22.5T416 128H96q-13 0-22.5 9.5T64 160v448q0 13 9.5 22.5T96 640h64q13 0 22.5-9.5T192 608V448h96q13 0 22.5-9.5T320 416v-64q0-13-9.5-22.5T288 320h-96v-64h224q13 0 22.5-9.5T448 224v-64zm507 0q8-13 3-22.5t-19-9.5h-67q-14 0-29 9.5T821 160l-53 99l-53-99q-7-13-22-22.5t-29-9.5h-67q-14 0-19 9.5t3 22.5l120 224l-120 224q-8 13-3 22.5t19 9.5h67q14 0 29-9.5t22-22.5l53-99l53 99q7 13 22 22.5t29 9.5h67q14 0 19-9.5t-3-22.5L835 384z"/>`),
		g.Group(children),
	)
}