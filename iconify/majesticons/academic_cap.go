package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AcademicCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11.514 3.126a1 1 0 0 1 .972 0l9 5A1 1 0 0 1 22 9v7a1 1 0 1 1-2 0v-5.3l-1 .555v.004l-6.067 3.016a2 2 0 0 1-1.848-.035L2.357 9.479a1 1 0 0 0-.284-.103a1 1 0 0 1 .441-1.25l9-5zM5 13.199V17a1 1 0 0 0 .553.894l6 3a1 1 0 0 0 .894 0l6-3A1 1 0 0 0 19 17v-3.256l-6.083 2.844a2 2 0 0 1-1.805-.056L5 13.2z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}