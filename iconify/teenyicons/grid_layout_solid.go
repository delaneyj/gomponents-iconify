package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridLayoutSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v4A1.5 1.5 0 0 0 1.5 7h4A1.5 1.5 0 0 0 7 5.5v-4A1.5 1.5 0 0 0 5.5 0h-4Zm8 0A1.5 1.5 0 0 0 8 1.5v4A1.5 1.5 0 0 0 9.5 7h4A1.5 1.5 0 0 0 15 5.5v-4A1.5 1.5 0 0 0 13.5 0h-4Zm-8 8A1.5 1.5 0 0 0 0 9.5v4A1.5 1.5 0 0 0 1.5 15h4A1.5 1.5 0 0 0 7 13.5v-4A1.5 1.5 0 0 0 5.5 8h-4Zm8 0A1.5 1.5 0 0 0 8 9.5v4A1.5 1.5 0 0 0 9.5 15h4a1.5 1.5 0 0 0 1.5-1.5v-4A1.5 1.5 0 0 0 13.5 8h-4Z"/>`),
		g.Group(children),
	)
}