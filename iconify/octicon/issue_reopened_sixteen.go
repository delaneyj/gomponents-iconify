package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueReopenedSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.029 2.217a6.5 6.5 0 0 1 9.437 5.11a.75.75 0 1 0 1.492-.154a8 8 0 0 0-14.315-4.03L.427 1.927A.25.25 0 0 0 0 2.104V5.75A.25.25 0 0 0 .25 6h3.646a.25.25 0 0 0 .177-.427L2.715 4.215a6.491 6.491 0 0 1 2.314-1.998ZM1.262 8.169a.75.75 0 0 0-1.22.658a8.001 8.001 0 0 0 14.315 4.03l1.216 1.216a.25.25 0 0 0 .427-.177V10.25a.25.25 0 0 0-.25-.25h-3.646a.25.25 0 0 0-.177.427l1.358 1.358a6.501 6.501 0 0 1-11.751-3.11a.75.75 0 0 0-.272-.506Z"/><path fill="currentColor" d="M9.06 9.06a1.5 1.5 0 1 1-2.12-2.12a1.5 1.5 0 0 1 2.12 2.12Z"/>`),
		g.Group(children),
	)
}