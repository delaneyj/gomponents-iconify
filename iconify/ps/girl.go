package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Girl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222v-11zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469zm0-128q21 0 21-21v-21h-42v21q0 21 21 21zm-64-106q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235zm171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235zM235 405h42q18 0 30.5-12.5T320 363H192q0 17 12.5 29.5T235 405z"/>`),
		g.Group(children),
	)
}