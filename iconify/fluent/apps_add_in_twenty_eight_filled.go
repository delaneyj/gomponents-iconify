package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsAddInTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 2a1 1 0 0 1 1 1v3h3a1 1 0 1 1 0 2h-3v3a1 1 0 1 1-2 0V8h-3a1 1 0 1 1 0-2h3V3a1 1 0 0 1 1-1ZM6 3a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3h-7V6a3 3 0 0 0-3-3H6Zm7 10H5V6a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v7Zm2 10v-8h7a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-7Zm-2 0H6a1 1 0 0 1-1-1v-7h8v8Z"/>`),
		g.Group(children),
	)
}