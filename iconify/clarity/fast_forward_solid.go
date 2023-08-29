package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M17.71 32a2 2 0 0 1-.86-.2A1.77 1.77 0 0 1 16 30v-6.7L5.17 31.58a1.94 1.94 0 0 1-2.06.22A2 2 0 0 1 2 30V6a2 2 0 0 1 1.11-1.8a1.93 1.93 0 0 1 2.06.22L16 12.69V6a1.77 1.77 0 0 1 .85-1.79a1.93 1.93 0 0 1 2.06.22l15.32 12a2 2 0 0 1 0 3.15l-15.32 12a2 2 0 0 1-1.2.42Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}