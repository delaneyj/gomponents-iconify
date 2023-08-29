package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M16 4a12 12 0 1 0 12 12A12 12 0 0 0 16 4Zm0 21.91A10 10 0 1 1 26 16a10 10 0 0 1-10 9.91Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="m31.71 29.69l-5.17-5.17A13.68 13.68 0 0 1 25.15 26l5.15 5.15a1 1 0 0 0 1.41-1.41Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M20 15h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}