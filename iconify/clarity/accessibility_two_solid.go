package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibilityTwoSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="17.96" cy="5" r="3" fill="currentColor" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M30 10H6a1 1 0 0 0 0 2h8v8.36l-3.89 12.81a1 1 0 0 0 .66 1.25a1.55 1.55 0 0 0 .29 0a1 1 0 0 0 1-.71l3.29-10.84h5.38L24 33.75a1 1 0 0 0 1 .71a1.55 1.55 0 0 0 .29 0a1 1 0 0 0 .66-1.25L22 20.4V12h8a1 1 0 0 0 0-2Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}