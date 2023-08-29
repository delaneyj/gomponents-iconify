package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlBigSmile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222v-11zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469zm0-149q21 0 21-21v-22h-42v22q0 21 21 21zm87 0q-20-3-23 19q0 4-3 13t-19 20.5t-42 11.5q-24 0-39.5-11.5t-20-21T192 337q-4-20-26-17q-8 2-13 9.5t-4 16.5q5 28 32 54.5t75 26.5q51 0 76.5-26.5T363 343q2-8-4-15t-16-8zm-2-64q9 0 15.5-6t6.5-15v-22h-43v22q0 9 6 15t15 6zm-170 0q8 0 14.5-6t6.5-15v-22h-43v22q0 9 6 15t16 6z"/>`),
		g.Group(children),
	)
}