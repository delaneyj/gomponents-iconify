package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M20 14h-4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1Zm-.4 6.6h-3.2v-5.2h3.2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="m33.71 12.38l-4.09-4.09a1 1 0 0 0-.7-.29h-5V6.05A2 2 0 0 0 22 4h-8.16A1.92 1.92 0 0 0 12 6.05V8H7.08a1 1 0 0 0-.71.29l-4.08 4.09a1 1 0 0 0-.29.71V28a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V13.08a1 1 0 0 0-.29-.7ZM14 6h8v2h-8Zm18 11H22v1.93h10V28H4v-9.07h10V17H4v-3.5L7.5 10h21l3.5 3.5Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}