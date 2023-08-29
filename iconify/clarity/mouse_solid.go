package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2A10 10 0 0 0 8 12v12a10 10 0 0 0 20 0V12A10 10 0 0 0 18 2Zm1.3 11.44a1.3 1.3 0 0 1-2.6 0V10a1.3 1.3 0 0 1 2.6 0Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}