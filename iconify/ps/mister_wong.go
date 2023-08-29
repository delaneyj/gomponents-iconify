package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MisterWong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M117 223q-18 19-56.5 52T17 313q-9-4-14-14q52-44 100-88q5 3 14 12zm182-5q-1 2-5.5 5.5T288 230q76 65 106 87q11-9 12-12q-16-12-51-42t-56-45zM99 77H3v17q17 1 48 1h49q15 115 101 135v113Q43 356 14 463q6 1 59 1q28-78 128-89v88q1 1 20 1q0-82 1-92q117 6 148 92h37q-27-105-186-122V229q77-20 94-134h92V77h-46.5L314 76q0-41-14-72h-32q16 47 12 73q-13-1-67.5 0T140 75q0-23 13-71h-39q-10 18-15 73zm45 18h136q-9 92-57 105q-14 3-26 0q-43-12-56-102q0-3 3-3z"/>`),
		g.Group(children),
	)
}