package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenterTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30.88 8H5.12a1.1 1.1 0 0 0 0 2.2h25.76a1.1 1.1 0 1 0 0-2.2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M25.5 16.2a1.1 1.1 0 1 0 0-2.2h-15a1.1 1.1 0 1 0 0 2.2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M30.25 20H5.75a1.1 1.1 0 0 0 0 2.2h24.5a1.1 1.1 0 0 0 0-2.2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M24.88 26H11.12a1.1 1.1 0 1 0 0 2.2h13.76a1.1 1.1 0 1 0 0-2.2Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}