package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsersSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5.5 0a3.499 3.499 0 1 0 0 6.996A3.499 3.499 0 1 0 5.5 0Zm-2 8.994a3.5 3.5 0 0 0-3.5 3.5v2.497h11v-2.497a3.5 3.5 0 0 0-3.5-3.5h-4Zm9 1.006H12v5h3v-2.5a2.5 2.5 0 0 0-2.5-2.5Z"/><path fill="currentColor" d="M11.5 4a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z"/>`),
		g.Group(children),
	)
}