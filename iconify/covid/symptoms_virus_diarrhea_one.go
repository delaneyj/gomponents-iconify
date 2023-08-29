package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymptomsVirusDiarrheaOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.886 10.392a3.506 3.506 0 1 0 0-7.012a3.506 3.506 0 0 0 0 7.012ZM6.302.75h1.169m-.585 0v2.63m3.926-1.246l.827.827m-.414-.414l-1.859 1.86m3.657 1.895v1.169m0-.585h-2.63m1.246 3.926l-.827.827m.413-.414L9.366 9.366m-1.895 3.657H6.302m.584 0v-2.63m-3.925 1.246l-.827-.827m.413.413l1.86-1.859M.75 7.471V6.302m0 .584h2.63M2.134 2.961l.827-.827m-.414.413l1.86 1.86m14.207 10.729h3.477a1.159 1.159 0 0 0 1.159-1.159V7.023a1.159 1.159 0 0 0-1.159-1.159h-2.318a1.159 1.159 0 0 0-1.159 1.159v8.113Z"/><path d="M20.509 17.335a2.32 2.32 0 0 0 1.582-2.2H9.341a6.374 6.374 0 0 0 3.477 5.68l-.579 2.434h7.534v-4.8a1.152 1.152 0 0 1 .736-1.114v0Z"/></g>`),
		g.Group(children),
	)
}