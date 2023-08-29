package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlOMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 203v-11q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zm-43 117q0 62-43.5 105.5T256 469t-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128zm-149 0q21 0 21-21v-22h-42v22q0 21 21 21zm-64-85q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235zm171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235zm-64 170q0 18-12.5 30.5T256 448t-30.5-12.5T213 405q0-17 12.5-29.5T256 363t30.5 12.5T299 405z"/>`),
		g.Group(children),
	)
}