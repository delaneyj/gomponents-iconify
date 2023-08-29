package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepForwardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M5 31.9a2 2 0 0 1-2-2V5.44a2 2 0 0 1 3.12-1.63L23.18 16a2 2 0 0 1 0 3.25L6.12 31.52A2 2 0 0 1 5 31.9Z" class="clr-i-solid clr-i-solid-path-1"/><rect width="7" height="28" x="25.95" y="3.67" fill="currentColor" class="clr-i-solid clr-i-solid-path-2" rx="2" ry="2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}