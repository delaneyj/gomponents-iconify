package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClosedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="17" r="1.3" fill="currentColor"/><path fill="currentColor" d="M17 10h-1V8c0-2.206-1.794-4-4-4S8 5.794 8 8v2H7a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2zm-7-2a2 2 0 0 1 4 0v3h-4V8zm7 11H7v-7h10.003L17 19z"/>`),
		g.Group(children),
	)
}