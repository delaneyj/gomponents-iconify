package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsReloadTwoArrowsLoadArrowSyncSquareLoadingReloadSynchronize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9 .5L11.5 3L9 5.5"/><path d="M.5 6V4a1 1 0 0 1 1-1h10M5 13.5L2.5 11L5 8.5"/><path d="M13.5 8v2a1 1 0 0 1-1 1h-10"/></g>`),
		g.Group(children),
	)
}