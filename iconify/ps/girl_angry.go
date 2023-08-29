package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlAngry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222v-11zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469zm0-128q21 0 21-21v-21h-42v21q0 21 21 21zm21 22h-42q-18 0-30.5 12.5T192 405q0-8 25-14.5t39-6.5t39 6.5t25 14.5q0-17-12.5-29.5T277 363zm-64-107v-43h-85l43 22v21q0 21 21 21t21-21zm86-43v43q0 21 21 21t21-21v-21l43-22h-85z"/>`),
		g.Group(children),
	)
}