package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M33 11H3a1 1 0 0 0 0 2h30a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M28 17H8a1 1 0 0 0 0 2h20a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M23 23H13a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}