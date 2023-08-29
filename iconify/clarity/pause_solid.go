package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<rect width="11" height="28" x="3.95" y="4" fill="currentColor" class="clr-i-solid clr-i-solid-path-1" rx="2.07" ry="2.07"/><rect width="11" height="28" x="20.95" y="4" fill="currentColor" class="clr-i-solid clr-i-solid-path-2" rx="2.07" ry="2.07"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}