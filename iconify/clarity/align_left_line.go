package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M5 1a1 1 0 0 0-1 1v32a1 1 0 0 0 2 0V2a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M31 20H8v10h23a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1Zm-1 8H10v-6h20Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M24 15V7a1 1 0 0 0-1-1H8v10h15a1 1 0 0 0 1-1Zm-2-1H10V8h12Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}