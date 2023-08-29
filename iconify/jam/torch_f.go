package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TorchF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 10h8v1a2 2 0 0 1-1.35 1.892l-.508 7.113a2.148 2.148 0 0 1-4.284 0l-.509-7.113A2 2 0 0 1 0 11v-1zm7.465-2H8H0h.535A3.982 3.982 0 0 1 0 6c0-1.473 1.333-3.473 4-6c2.667 2.527 4 4.527 4 6c0 .729-.195 1.412-.535 2z"/>`),
		g.Group(children),
	)
}