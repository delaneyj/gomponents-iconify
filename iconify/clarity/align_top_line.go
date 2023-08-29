package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTopLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 4H2a1 1 0 0 0 0 2h32a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M6 31a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V8H6Zm2-21h6v20H8Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M20 23a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V8H20Zm2-13h6v12h-6Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}