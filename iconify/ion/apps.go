package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M104 160a56 56 0 1 1 56-56a56.06 56.06 0 0 1-56 56Zm152 0a56 56 0 1 1 56-56a56.06 56.06 0 0 1-56 56Zm152 0a56 56 0 1 1 56-56a56.06 56.06 0 0 1-56 56ZM104 312a56 56 0 1 1 56-56a56.06 56.06 0 0 1-56 56Zm152 0a56 56 0 1 1 56-56a56.06 56.06 0 0 1-56 56Zm152 0a56 56 0 1 1 56-56a56.06 56.06 0 0 1-56 56ZM104 464a56 56 0 1 1 56-56a56.06 56.06 0 0 1-56 56Zm152 0a56 56 0 1 1 56-56a56.06 56.06 0 0 1-56 56Zm152 0a56 56 0 1 1 56-56a56.06 56.06 0 0 1-56 56Z"/>`),
		g.Group(children),
	)
}