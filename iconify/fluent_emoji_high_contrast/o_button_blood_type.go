package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OButtonBloodType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M16.008 7a5.977 5.977 0 0 0-5.977 5.977v6.046a5.977 5.977 0 1 0 11.953 0v-6.046A5.976 5.976 0 0 0 16.008 7Zm-2.977 5.977a2.977 2.977 0 0 1 5.953 0v6.046a2.977 2.977 0 0 1-5.953 0v-6.046Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}