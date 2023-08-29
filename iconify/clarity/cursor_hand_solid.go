package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorHandSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M28.69 14.33v4.83l-2-.43v-5.49a16.19 16.19 0 0 0-2.33-.84v5.82l-2-.43V12c-1.1-.18-2.18-.3-3.08-.36v5.51l-2-.43V4.34a2.53 2.53 0 0 0-2.6-2.43a2.53 2.53 0 0 0-2.6 2.43v15.52l-2 1V15.6l-2.33-2.39a2.83 2.83 0 0 0-4 0a2.93 2.93 0 0 0 0 4.09l6 7.1a10.82 10.82 0 0 0 1.39 4.22a8.42 8.42 0 0 0 2.21 2.73v2.56h14.44v-3.29a12.54 12.54 0 0 0 3-8.5v-6a10 10 0 0 0-2.1-1.79Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}