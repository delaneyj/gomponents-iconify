package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 107h-23v-7q21-10 21-36V43q0-18-12.5-30.5T425 0H85Q68 0 55.5 12.5T43 43v21q0 26 21 36v7H43q-18 0-30.5 12.5T0 149q0 53 37.5 90.5T128 277h41q11 10 23 15v94q0 54-21 75q-9 8-22 8h-21q-21 0-21 22q0 21 21 21h258q21 0 21-21q0-22-21-22h-21q-17 0-24-8q-21-19-21-75v-94q16-7 23-15h41q53 0 90.5-37.5T512 149q0-17-12.5-29.5T469 107zM85 43h342v21H85V43zM43 149h34q13 40 47 86q-34-1-57.5-26.5T43 149zm172 320q11-19 15-42h49q8 32 13 42h-77zm62-85h-42v-77q6 2 21 2t21-2v77zm22-128h-86q-41-20-73.5-64.5T107 107h298q0 39-33 84t-73 65zm89-21q30-36 47-86h34q0 34-23.5 59.5T388 235z"/>`),
		g.Group(children),
	)
}