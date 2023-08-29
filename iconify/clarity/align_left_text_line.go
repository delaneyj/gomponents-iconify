package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M20.25 26H6v2.2h14.25a1.1 1.1 0 0 0 0-2.2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M28 20H6v2.2h22a1.1 1.1 0 0 0 0-2.2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M22.6 15.1a1.1 1.1 0 0 0-1.1-1.1H6v2.2h15.5a1.1 1.1 0 0 0 1.1-1.1Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M29.25 8H6v2.2h23.25a1.1 1.1 0 1 0 0-2.2Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}