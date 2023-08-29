package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PasteLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 12h-4v2h4v2h2v-2a2 2 0 0 0-2-2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M30 18h2v6h-2z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M30 30h-2v2h2a2 2 0 0 0 2-2v-4h-2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M24 22V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2ZM6 6h16v16H6Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M20 30h6v2h-6z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M14 26h-2v4a2 2 0 0 0 2 2h4v-2h-4Z" class="clr-i-outline clr-i-outline-path-6"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}