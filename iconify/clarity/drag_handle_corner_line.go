package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragHandleCornerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="12" cy="24" r="1.5" fill="currentColor" class="clr-i-outline clr-i-outline-path-1"/><circle cx="18" cy="24" r="1.5" fill="currentColor" class="clr-i-outline clr-i-outline-path-2"/><circle cx="18" cy="18" r="1.5" fill="currentColor" class="clr-i-outline clr-i-outline-path-3"/><circle cx="24" cy="12" r="1.5" fill="currentColor" class="clr-i-outline clr-i-outline-path-4"/><circle cx="24" cy="24" r="1.5" fill="currentColor" class="clr-i-outline clr-i-outline-path-5"/><circle cx="24" cy="18" r="1.5" fill="currentColor" class="clr-i-outline clr-i-outline-path-6"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}