package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackspaceSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.11 2.188A.5.5 0 0 1 4.5 2h9A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-9a.5.5 0 0 1-.39-.188l-4-5a.5.5 0 0 1 0-.624l4-5Zm6.036 7.666L8.5 8.207L6.854 9.854l-.708-.708L7.793 7.5L6.146 5.854l.708-.708L8.5 6.793l1.646-1.647l.708.708L9.207 7.5l1.647 1.646l-.707.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}