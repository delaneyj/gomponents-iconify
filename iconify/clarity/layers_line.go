package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 20.25a1 1 0 0 1-.43-.1l-15-7.09a1 1 0 0 1 0-1.81l15-7.09a1 1 0 0 1 .85 0l15 7.09a1 1 0 0 1 0 1.81l-15 7.09a1 1 0 0 1-.42.1ZM5.34 12.16l12.66 6l12.66-6L18 6.18Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 26.16a1 1 0 0 1-.43-.1L2.57 19a1 1 0 1 1 .85-1.81L18 24.06l14.57-6.89a1 1 0 1 1 .86 1.83l-15 7.09a1 1 0 0 1-.43.07Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M18 32.07a1 1 0 0 1-.43-.1l-15-7.09a1 1 0 0 1 .85-1.81L18 30l14.57-6.89a1 1 0 1 1 .85 1.81L18.43 32a1 1 0 0 1-.43.07Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}