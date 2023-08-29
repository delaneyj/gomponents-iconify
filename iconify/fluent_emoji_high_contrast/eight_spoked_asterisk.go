package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightSpokedAsterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15 7a1 1 0 1 1 2 0v6.586l4.657-4.657a1 1 0 0 1 1.414 1.414L18.414 15H25a1 1 0 1 1 0 2h-6.586l4.657 4.657a1 1 0 0 1-1.414 1.414L17 18.414V25a1 1 0 1 1-2 0v-6.586l-4.657 4.657a1 1 0 0 1-1.414-1.414L13.586 17H7a1 1 0 1 1 0-2h6.586l-4.657-4.657a1 1 0 0 1 1.414-1.414L15 13.586V7Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}