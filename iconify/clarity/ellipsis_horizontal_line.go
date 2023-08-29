package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisHorizontalLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="31.1" cy="18" r="2.9" fill="currentColor" class="clr-i-outline clr-i-outline-path-1"/><circle cx="18" cy="18" r="2.9" fill="currentColor" class="clr-i-outline clr-i-outline-path-2"/><circle cx="4.9" cy="18" r="2.9" fill="currentColor" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}