package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoundLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M27.9 30H13.4a8.45 8.45 0 0 0 1.6-5.35V21h4.31a1 1 0 0 0 0-2H15v-7.69A5.24 5.24 0 0 1 20.21 6A5.19 5.19 0 0 1 24 7.73a1 1 0 0 0 1.48-1.35A7.19 7.19 0 0 0 13 11.31V19H8.72a1 1 0 1 0 0 2H13v3.65c0 4.73-2.88 5.35-3 5.35a1 1 0 0 0 .17 2H27.9a1 1 0 1 0 0-2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}