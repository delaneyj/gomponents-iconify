package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cigarette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 12H2v4h16m4-4v4M7 12v4m11-8c0-2.5-2-2.5-2-5m6 5c0-2.5-2-2.5-2-5"/>`),
		g.Group(children),
	)
}