package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M214 193.68L145.35 128L214 62.32l.18-.18A20 20 0 0 0 200 28H56a20 20 0 0 0-14.13 34.14l.18.18l68.6 65.68l-68.6 65.68l-.18.18A20 20 0 0 0 56 228h144a20 20 0 0 0 14.14-34.14ZM190 52l-62 59.39L66 52ZM66 204l62-59.39L190 204Z"/>`),
		g.Group(children),
	)
}