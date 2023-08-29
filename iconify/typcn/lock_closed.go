package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 10h-1V8c0-2.205-1.794-4-4-4S8 5.795 8 8v2H7c-1.103 0-2 .896-2 2v7c0 1.104.897 2 2 2h10c1.103 0 2-.896 2-2v-7c0-1.104-.897-2-2-2zm-5 8.299a1.3 1.3 0 1 1 0-2.6a1.3 1.3 0 0 1 0 2.6zM14 11h-4V8c0-1.104.897-2 2-2s2 .896 2 2v3z"/>`),
		g.Group(children),
	)
}