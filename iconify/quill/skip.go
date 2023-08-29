package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 5v22M8 7.586v16.828c0 1.746 2.081 2.653 3.36 1.465l9.062-8.413a2 2 0 0 0 0-2.932L11.36 6.121C10.08 4.933 8 5.84 8 7.586Z"/>`),
		g.Group(children),
	)
}