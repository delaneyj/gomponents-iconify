package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M28 2H8a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2ZM12 28h-2v-2h2Zm0-6h-2v-2h2Zm0-6h-2v-2h2Zm7 12h-2v-2h2Zm0-6h-2v-2h2Zm0-6h-2v-2h2Zm7 12h-2v-2h2Zm0-6h-2v-2h2Zm0-6h-2v-2h2Zm0-7H10V5h16Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}