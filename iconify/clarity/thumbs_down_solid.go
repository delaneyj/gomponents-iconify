package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M16.37 23.84c2.12 2.84 4.76 10.07 4.76 10.07s2.87-.78 2.87-3.2V21h9.77a29.46 29.46 0 0 0-1.44-9.74C30.39 5.68 28.2 4 28.2 4h-6.85C19.1 4 15 5.9 12 7.65v12.8a10.84 10.84 0 0 1 4.37 3.39Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M9 23a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H2v18Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}