package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GuyBigSmile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171zM64 213v69q-21-18-21-47q0-22 21-22zm192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469zm192-187v-69q21 0 21 22q0 29-21 47zm-192 38q21 0 21-21v-22h-42v22q0 21 21 21zm87 0q-20-3-23 19q0 4-3 13t-19 20.5t-42 11.5q-24 0-39.5-11.5t-20-21T192 337q-4-20-26-17q-8 2-13 9.5t-4 16.5q5 28 32 54.5t75 26.5q51 0 76.5-26.5T363 343q2-8-4-15t-16-8zm-2-64q9 0 15.5-6t6.5-15v-22h-43v22q0 9 6 15t15 6zm-170 0q8 0 14.5-6t6.5-15v-22h-43v22q0 9 6 15t16 6z"/>`),
		g.Group(children),
	)
}