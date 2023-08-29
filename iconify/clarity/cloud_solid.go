package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M29 16.66a10.07 10.07 0 0 0 .25-2.24a10.19 10.19 0 0 0-20.34-1.06A10 10 0 0 0 1 23.1c0 5.09 4.62 9.9 9.57 9.9h16.52c4.19 0 7.91-3.9 7.91-8.35a8.29 8.29 0 0 0-6-7.99Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}