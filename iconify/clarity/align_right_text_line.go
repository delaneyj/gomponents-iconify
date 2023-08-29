package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M14.65 27.1a1.1 1.1 0 0 0 1.1 1.1H30V26H15.75a1.1 1.1 0 0 0-1.1 1.1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M6.9 21.1A1.1 1.1 0 0 0 8 22.2h22V20H8a1.1 1.1 0 0 0-1.1 1.1Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M13.4 15.1a1.1 1.1 0 0 0 1.1 1.1H30V14H14.5a1.1 1.1 0 0 0-1.1 1.1Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M6.75 8a1.1 1.1 0 1 0 0 2.2H30V8Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}