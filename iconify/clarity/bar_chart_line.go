package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 5H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2ZM4 29V7h28v22Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M7 10h6v16h-1.6V11.6H8.6V26H7Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M15 19h6v7h-1.6v-5.4h-2.8V26H15Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M23 16h6v10h-1.6v-8.4h-2.8V26H23Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}