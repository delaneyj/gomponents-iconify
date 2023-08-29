package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 4c-2.206 0-4 1.795-4 4v3h-4v-1H7c-1.103 0-2 .896-2 2v7c0 1.104.897 2 2 2h10c1.103 0 2-.896 2-2v-7c0-1.104-.897-2-2-2h-1V8c0-1.104.897-2 2-2s2 .896 2 2v3a1 1 0 1 0 2 0V8c0-2.205-1.794-4-4-4zm-6 14.299a1.3 1.3 0 1 1 0-2.6a1.3 1.3 0 0 1 0 2.6z"/>`),
		g.Group(children),
	)
}