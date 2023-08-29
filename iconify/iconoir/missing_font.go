package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MissingFont(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.469 18.374l1.064-2.341m9.58 2.341l-1.064-2.341m0 0L8.79 6.664l-4.258 9.367m8.516 0H4.533m10.645-7.237c0-3.725 5.854-3.725 5.854 0c0 2.661-2.66 2.13-2.66 5.322m-.001 4.269l.01-.012"/>`),
		g.Group(children),
	)
}