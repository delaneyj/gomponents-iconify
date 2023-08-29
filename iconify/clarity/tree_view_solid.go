package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeViewSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<rect width="6" height="6" x="10" y="26" fill="currentColor" class="clr-i-solid clr-i-solid-path-1" rx="1" ry="1"/><path fill="currentColor" d="M15 16h-4a1 1 0 0 0-1 1v1.2H5.8V12H7a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1.2v17.8H11a.8.8 0 1 0 0-1.6H5.8v-8.4H10V21a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M33 8H10v2h23a1 1 0 0 0 0-2Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="currentColor" d="M33 18H18v2h15a1 1 0 0 0 0-2Z" class="clr-i-solid clr-i-solid-path-4"/><path fill="currentColor" d="M33 28H18v2h15a1 1 0 0 0 0-2Z" class="clr-i-solid clr-i-solid-path-5"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}