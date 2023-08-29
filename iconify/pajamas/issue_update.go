package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueUpdate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3.25a.75.75 0 0 0 0-1.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 .5.5v5.25a.75.75 0 0 0 1.5 0V4.364l2.19 1.14a.25.25 0 0 1 .107.338l-1.072 2.062a.75.75 0 0 0 1.33.692l1.073-2.062a1.75 1.75 0 0 0-.745-2.36l-2.912-1.516A2 2 0 0 0 9 1H3Zm10 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}