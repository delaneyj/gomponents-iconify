package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PopOutLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M27 33H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h10v2H5v22h22V21h2v10a2 2 0 0 1-2 2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 3a1 1 0 0 0 0 2h11.59L15.74 18.85a1 1 0 1 0 1.41 1.41L31 6.41V18a1 1 0 0 0 2 0V3Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}