package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unapproval(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM7 5a2.99 2.99 0 0 1-.87 2.113A3.997 3.997 0 0 1 8 10.5V12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-1.5c0-1.427.747-2.679 1.87-3.387A3 3 0 1 1 7 5Zm-5.5 5.5a2.5 2.5 0 0 1 5 0V12a.5.5 0 0 1-.5.5H2a.5.5 0 0 1-.5-.5v-1.5Zm9.72-6.28a.75.75 0 0 1 1.06 0l1.22 1.22l1.22-1.22a.75.75 0 1 1 1.06 1.06L14.56 6.5l1.22 1.22a.75.75 0 0 1-1.06 1.06L13.5 7.56l-1.22 1.22a.75.75 0 1 1-1.06-1.06l1.22-1.22l-1.22-1.22a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}