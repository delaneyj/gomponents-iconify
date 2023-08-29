package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrollSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 11.12V6.58a4.5 4.5 0 0 0-4.5-4.5h-16A4.5 4.5 0 0 0 9 6.58V24H2v5.5A4.5 4.5 0 0 0 6.5 34h19.08a4.5 4.5 0 0 0 4.5-4.5V13.13h-2v16.41a2.5 2.5 0 0 1-2.5 2.5H10.24a4.47 4.47 0 0 0 .76-2.5v-23a2.5 2.5 0 0 1 5 0v4.54Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}