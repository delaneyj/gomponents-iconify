package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GuyAngry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171zM64 213v69q-21-18-21-47q0-22 21-22zm192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469zm192-187v-69q21 0 21 22q0 29-21 47zm-192 59q21 0 21-21v-21h-42v21q0 21 21 21zm21 22h-42q-18 0-30.5 12.5T192 405q0-8 25-14.5t39-6.5t39 6.5t25 14.5q0-17-12.5-29.5T277 363zm-64-107v-43h-85l43 22v21q0 21 21 21t21-21zm86-43v43q0 21 21 21t21-21v-21l43-22h-85z"/>`),
		g.Group(children),
	)
}