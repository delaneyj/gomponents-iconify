package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M5 17.5h2m2 0h6m2 0h2M2.5 4.5h19m0 9s-3.969 1-9.5 1c-5.531 0-9.5-1-9.5-1m0-12h19v21h-2.25l-.075-.12a4 4 0 0 0-3.392-1.88H8.217a4 4 0 0 0-3.392 1.88l-.075.12H2.5v-21Z"/>`),
		g.Group(children),
	)
}