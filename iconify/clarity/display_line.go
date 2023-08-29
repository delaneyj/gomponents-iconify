package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplayLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32.5 3h-29A1.5 1.5 0 0 0 2 4.5v21A1.5 1.5 0 0 0 3.5 27h29a1.5 1.5 0 0 0 1.5-1.5v-21A1.5 1.5 0 0 0 32.5 3ZM32 25H4V5h28Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M7.7 8.76h20.43l1.81-1.6H6.1V23h1.6V8.76z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M26 32h-1.74a3.61 3.61 0 0 1-1.5-2.52v-1.35h-1.52v1.37a4.2 4.2 0 0 0 .93 2.5h-8.34a4.2 4.2 0 0 0 .93-2.52v-1.35h-1.52v1.37a3.61 3.61 0 0 1-1.5 2.5h-1.8a1 1 0 1 0 0 2h16.12a.92.92 0 0 0 1-1A1 1 0 0 0 26 32Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}