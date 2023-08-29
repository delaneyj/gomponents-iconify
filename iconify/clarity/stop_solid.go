package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<rect width="27.99" height="28" x="3.96" y="4" fill="currentColor" class="clr-i-solid clr-i-solid-path-1" rx="2" ry="2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}