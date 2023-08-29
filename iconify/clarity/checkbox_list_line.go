package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxListLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.43 16H10v2h21.43a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M31.43 24H10v2h21.43a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M15.45 10h16a1 1 0 0 0 0-2h-14Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M17.5 3.42a1.09 1.09 0 0 0-1.55 0l-8.06 8.06l-3.38-3.64a1.1 1.1 0 1 0-1.61 1.5l4.94 5.3L17.5 5a1.1 1.1 0 0 0 0-1.58Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}