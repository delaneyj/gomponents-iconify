package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDocumentArrowRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23 19l-3-3v2h-4v2h4v2l3-3m-9.2 3H6c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h8l6 6v5.1c-.3-.1-.7-.1-1-.1s-.7 0-1 .1V9h-5V4H6v16h7.1c.1.7.4 1.4.7 2M8 12h8v1.8c-.1.1-.2.1-.3.2H8v-2m0 4h5v2H8v-2Z"/>`),
		g.Group(children),
	)
}