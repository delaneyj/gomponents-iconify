package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminMultisite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.27 6.87L10 3.14L5.73 6.87L5 6.14l5-4.38l5 4.38zM14 8.42l-4.05 3.43L6 8.38v-.74l4-3.5l4 3.5v.78zM11 9.7V8H9v1.7h2zm-1.73 4.03L5 10L.73 13.73L0 13l5-4.38L10 13zm10 0L15 10l-4.27 3.73L10 13l5-4.38L20 13zM5 11l4 3.5V18H1v-3.5zm10 0l4 3.5V18h-8v-3.5zm-9 6v-2H4v2h2zm10 0v-2h-2v2h2z"/>`),
		g.Group(children),
	)
}