package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M27.36 4H6a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V8.78ZM25 30H11v-8h14Zm5 0h-3v-8a2 2 0 0 0-2-2H11a2 2 0 0 0-2 2v8H6V6h4v6a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2H12V6h14.51L30 9.59Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}