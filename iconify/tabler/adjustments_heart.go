package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustmentsHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10a2 2 0 1 0 4 0a2 2 0 0 0-4 0m2-6v4m0 4v8m6-16v8.5M16 7a2 2 0 1 0 4 0a2 2 0 0 0-4 0m2-3v1m0 4v2.5M18 22l3.35-3.284a2.143 2.143 0 0 0 .005-3.071a2.242 2.242 0 0 0-3.129-.006l-.224.22l-.223-.22a2.242 2.242 0 0 0-3.128-.006a2.143 2.143 0 0 0-.006 3.071L18 22z"/>`),
		g.Group(children),
	)
}