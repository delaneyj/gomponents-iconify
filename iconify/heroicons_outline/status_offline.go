package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusOffline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636a9 9 0 0 1 0 12.728m0 0l-2.829-2.829m2.829 2.829L21 21M15.536 8.464a5 5 0 0 1 0 7.072m0 0l-2.829-2.829m-4.243 2.829a4.978 4.978 0 0 1-1.414-2.83m-1.414 5.658a9 9 0 0 1-2.167-9.238m7.824 2.167a1 1 0 1 1 1.414 1.414m-1.414-1.414L3 3m8.293 8.293l1.414 1.414"/>`),
		g.Group(children),
	)
}