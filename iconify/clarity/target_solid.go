package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="18" r="4.09" fill="currentColor" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M18 7.83A10.17 10.17 0 1 0 28.17 18A10.18 10.18 0 0 0 18 7.83Zm0 16A5.88 5.88 0 1 1 23.88 18A5.88 5.88 0 0 1 18 23.88Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2Zm0 27.83A11.83 11.83 0 1 1 29.83 18A11.85 11.85 0 0 1 18 29.83Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}