package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsReloadOneArrowsLoadArrowSyncSquareLoadingReloadSynchronize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 5L3 2.5L5.5 5"/><path d="M6 13.5H4a1 1 0 0 1-1-1v-10M13.5 9L11 11.5L8.5 9"/><path d="M8 .5h2a1 1 0 0 1 1 1v10"/></g>`),
		g.Group(children),
	)
}