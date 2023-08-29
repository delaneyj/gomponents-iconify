package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M29.2 26.72a12.07 12.07 0 0 1-6.3-22.28A13.68 13.68 0 0 0 19.49 4a14 14 0 0 0 0 28a13.82 13.82 0 0 0 10.9-5.34a11.71 11.71 0 0 1-1.19.06Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}