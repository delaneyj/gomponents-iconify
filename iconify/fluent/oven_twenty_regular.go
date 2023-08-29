package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvenTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 3a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6ZM4 14V8h12v6a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2Zm12-7H4V6a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v1Zm-9.25-.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm6.5 0a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm-2.5-.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM6 9a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1H6Zm0 5v-4h8v4H6Z"/>`),
		g.Group(children),
	)
}