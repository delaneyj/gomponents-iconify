package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintRollerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<rect width="27" height="10" x="4" y="2" fill="currentColor" class="clr-i-solid clr-i-solid-path-1" rx="1" ry="1"/><path fill="currentColor" d="M33 6h-1v6.24l-13.29 4.21a1 1 0 0 0-.71 1V19h-2v15a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V19h-2v-.82L33.29 14a1 1 0 0 0 .71-1V7a1 1 0 0 0-1-1Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}