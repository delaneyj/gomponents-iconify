package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloneLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M6 6h16v4h2V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h4v-2H6Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M30 12H14a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V14a2 2 0 0 0-2-2Zm0 18H14V14h16Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M21 28h2v-5h5v-2h-5v-5h-2v5h-5v2h5v5z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}