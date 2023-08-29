package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortByLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M28.54 13H7.46a1 1 0 0 1 0-2h21.08a1 1 0 0 1 0 2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M21.17 19H7.46a1 1 0 0 1 0-2h13.71a1 1 0 0 1 0 2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M13.74 25H7.46a1 1 0 0 1 0-2h6.28a1 1 0 0 1 0 2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}