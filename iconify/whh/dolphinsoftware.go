package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dolphinsoftware(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M883 159q67 70 104.5 161t37.5 192q0 146-78 270q30 40 78 44q-33 38-82 38q-44 0-77-40t-33-88q0-20 5-64.5t6-65.5q-33-73-105.5-128.5T580 398q-35 15-78 49t-77 66t-77.5 63t-80.5 42q-99 29-176.5 19T2 583q-7-31 23.5-67t88.5-68q-2-2-4.5-6.5T106 436q53 12 119 12q53 0 97-17.5t69-47.5q25-2 41.5-20t16.5-43q0-26-18.5-45t-45-19t-45.5 19t-19 45q0 22 15 40q-51 24-111 24q-38 0-59-1t-46.5-5.5T76 364q-10-41-11-76q-1-49 14-87t50-73q64-63 169-95.5T513 0q129 0 242 61Q809 0 888 0q82 0 137 66q-45 4-82 29t-60 64zM180 702q50 88 139 141t194 53q147 0 257-99q7 68 58 118q-139 109-315 109q-159 0-288-89T38 703q27 1 59 1q48 0 83-2z"/>`),
		g.Group(children),
	)
}