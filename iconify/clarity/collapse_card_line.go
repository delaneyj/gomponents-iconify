package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseCardLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M33 21H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h30a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1Zm-1 6H4v-4h28Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="m18 20.22l5.65-5.65a.81.81 0 0 0 0-1.14a.8.8 0 0 0-1.13 0L18 18l-4.52-4.52a.8.8 0 0 0-1.13 0a.81.81 0 0 0 0 1.14Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="m18 14.22l5.65-5.65a.81.81 0 0 0 0-1.14a.8.8 0 0 0-1.13 0L18 12l-4.52-4.57a.8.8 0 0 0-1.13 0a.81.81 0 0 0 0 1.14Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}