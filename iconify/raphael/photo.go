package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.25 10.25H20.5v-1.5h-9.375v1.5h-3.75a2 2 0 0 0-2 2v10.375a2 2 0 0 0 2 2H24.25a2 2 0 0 0 2-2V12.25a2 2 0 0 0-2-2zM15.812 23.5c-3.342 0-6.06-2.72-6.06-6.062s2.718-6.062 6.06-6.062s6.062 2.72 6.062 6.062s-2.72 6.06-6.062 6.06zm0-10.125a4.062 4.062 0 1 0 .001 8.127a4.062 4.062 0 0 0-.001-8.128z"/>`),
		g.Group(children),
	)
}