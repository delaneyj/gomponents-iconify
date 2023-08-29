package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M5 6a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v36a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6Z"/><circle cx="11" cy="35" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M11 11v14M23.977 9.756l10.06-2.515a1 1 0 0 1 1.211.72l7.5 29.063a1 1 0 0 1-.725 1.22l-10.06 2.515a1 1 0 0 1-1.211-.72l-7.5-29.063a1 1 0 0 1 .726-1.22Z"/><circle cx="35" cy="32" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m31 16l2.5 10"/></g>`),
		g.Group(children),
	)
}