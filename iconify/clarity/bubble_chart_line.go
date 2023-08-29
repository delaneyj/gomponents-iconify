package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 5H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2ZM4 29V7h28v22Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M29 18a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-3-1.4a1.4 1.4 0 1 0 0 2.8a1.4 1.4 0 0 0 0-2.8Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M15 14a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-4-2.4a2.4 2.4 0 1 0 .001 4.801A2.4 2.4 0 0 0 11 11.6Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M21 23a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-3-1.4a1.4 1.4 0 1 0 0 2.8a1.4 1.4 0 0 0 0-2.8Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}