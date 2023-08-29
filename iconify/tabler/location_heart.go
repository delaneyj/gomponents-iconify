package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.365 14.73L10 14l-7-3.5a.55.55 0 0 1 0-1L21 3l-3.024 8.373M18 22l3.35-3.284a2.143 2.143 0 0 0 .005-3.071a2.242 2.242 0 0 0-3.129-.006l-.224.22l-.223-.22a2.242 2.242 0 0 0-3.128-.006a2.143 2.143 0 0 0-.006 3.071L18 22z"/>`),
		g.Group(children),
	)
}