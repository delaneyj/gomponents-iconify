package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M256 128q0 72-64 110v1266q0 13-9.5 22.5T160 1536H96q-13 0-22.5-9.5T64 1504V238Q0 200 0 128q0-53 37.5-90.5T128 0t90.5 37.5T256 128zm1472 64v763q0 25-12.5 38.5T1676 1021q-215 116-369 116q-61 0-123.5-22t-108.5-48t-115.5-48T817 997q-192 0-464 146q-17 9-33 9q-26 0-45-19t-19-45V346q0-32 31-55q21-14 79-43q236-120 421-120q107 0 200 29t219 88q38 19 88 19q54 0 117.5-21t110-47t88-47t54.5-21q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}