package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Booth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 5h40v8l-1.398.84a7 7 0 0 1-7.203 0L34 13l-1.398.84a7 7 0 0 1-7.203 0L24 13l-1.398.84a7 7 0 0 1-7.204 0L14 13l-1.399.84a7 7 0 0 1-7.202 0L4 13V5Zm2 20h36v18H6zm3-9v9m30-9v9"/>`),
		g.Group(children),
	)
}