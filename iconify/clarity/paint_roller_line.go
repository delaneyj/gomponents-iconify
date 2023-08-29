package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintRollerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31 10V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h23a2 2 0 0 0 2-2ZM6 4h23v6H6Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M33 6h-1v6.29l-13.3 4.25a1 1 0 0 0-.7 1V19h-2v14a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V19h-2v-.73L33.3 14a1 1 0 0 0 .7-1V7a1 1 0 0 0-1-1ZM20 33h-2V21h2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}