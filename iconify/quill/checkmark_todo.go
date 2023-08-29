package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkTodo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M26.207 9.293a1 1 0 0 1 0 1.414l-1.625 1.625a1 1 0 0 1-1.414-1.414l1.625-1.625a1 1 0 0 1 1.414 0Zm-4.875 4.875a1 1 0 0 1 0 1.414l-3.25 3.25a1 1 0 0 1-1.414-1.414l3.25-3.25a1 1 0 0 1 1.414 0Zm-15.54 2.125a1 1 0 0 1 1.415 0l1.5 1.5a1 1 0 1 1-1.414 1.414l-1.5-1.5a1 1 0 0 1 0-1.414Zm9.04 4.375a1 1 0 0 1 0 1.414l-1.625 1.625a1 1 0 0 1-1.414 0l-1.5-1.5a1 1 0 1 1 1.414-1.414l.793.793l.918-.918a1 1 0 0 1 1.414 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}