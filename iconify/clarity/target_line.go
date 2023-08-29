package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2Zm0 30a14 14 0 1 1 14-14a14 14 0 0 1-14 14Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 7.2A10.8 10.8 0 1 0 28.8 18A10.81 10.81 0 0 0 18 7.2Zm0 20a9.2 9.2 0 1 1 9.2-9.2a9.21 9.21 0 0 1-9.2 9.2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M18 12.31A5.69 5.69 0 1 0 23.69 18A5.69 5.69 0 0 0 18 12.31Zm0 9.77A4.09 4.09 0 1 1 22.09 18A4.09 4.09 0 0 1 18 22.09Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}