package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlCry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222v-11zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469zm0-149q21 0 21-21v-22h-42v22q0 21 21 21zm85-107q-8 0-14.5 6.5T320 235q0 8 6 14.5t15 6.5t15.5-6t6.5-15t-6.5-15.5T341 213zm-21 107q0 8 6 14.5t15 6.5q22 0 22-21v-21h-43v21zm-149-64q8 0 14.5-6t6.5-15t-6-15.5t-15-6.5t-15.5 6.5T149 235t6.5 15t15.5 6zm21 64v-21h-43v21q0 8 6.5 14.5T171 341q21 0 21-21zm64 43q-50 0-79 30q-6 6-6 14.5t6 14.5q15 15 30 0q19-17 49-17q27 0 49 17q6 7 15 7t15-7q6-6 6-14.5t-6-14.5q-29-30-79-30z"/>`),
		g.Group(children),
	)
}