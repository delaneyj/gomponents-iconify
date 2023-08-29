package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTypeMaintenance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 0a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4zm7.293 3.292a3 3 0 0 0-4.001 4.001l-4 4a1 1 0 1 0 1.415 1.414l4-4a3 3 0 0 0 4.001-4.001l-2 2.001a1 1 0 1 1-1.415-1.414l2-2z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}