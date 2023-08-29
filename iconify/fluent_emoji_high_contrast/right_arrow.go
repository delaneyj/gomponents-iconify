package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M24.833 15.621a1 1 0 0 1 0 1.387l-5.112 5.313c-.625.649-1.721.207-1.721-.693v-3.063a.25.25 0 0 0-.25-.25H8a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h9.75a.25.25 0 0 0 .25-.25v-3.063c0-.9 1.096-1.342 1.72-.693l5.113 5.312Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}