package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JustifyTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M6 10.2h25.75a1.1 1.1 0 1 0 0-2.2H6a1.1 1.1 0 1 0 0 2.2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M31.75 14H6a1.1 1.1 0 1 0 0 2.2h25.75a1.1 1.1 0 1 0 0-2.2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M31.12 20H6.62a1.1 1.1 0 1 0 0 2.2h24.5a1.1 1.1 0 1 0 0-2.2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M30.45 25.83H6.6a1.1 1.1 0 0 0 0 2.2h23.85a1.1 1.1 0 0 0 0-2.2Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}