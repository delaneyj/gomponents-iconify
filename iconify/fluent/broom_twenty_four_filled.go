package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BroomTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22.453 1.923a.75.75 0 0 1 0 1.06L16.14 9.296a6.814 6.814 0 0 0-1.068-1.053l6.32-6.32a.75.75 0 0 1 1.061 0ZM15.03 9.347a5.75 5.75 0 0 0-8.132 0l-.244.244l8.132 8.132l.244-.244a5.75 5.75 0 0 0 0-8.132ZM1.885 12.966l3.613-2.409l8.32 8.321l-2.408 3.613a.75.75 0 0 1-1.154.115L1.77 14.12a.75.75 0 0 1 .115-1.154Z"/>`),
		g.Group(children),
	)
}