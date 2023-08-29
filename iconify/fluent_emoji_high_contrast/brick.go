package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 2a3 3 0 0 0-3 3v2h4a1 1 0 0 0 1-1V2H5Zm4 0v4a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2H9Zm16 0v4a1 1 0 0 0 1 1h4V5a3 3 0 0 0-3-3h-2Zm5 7H18a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h12V9Zm0 8h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4v-6Zm0 8H18a1 1 0 0 0-1 1v4h10a3 3 0 0 0 3-3v-2Zm-15 5v-4a1 1 0 0 0-1-1H2v2a3 3 0 0 0 3 3h10ZM2 23h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1H2v6Zm0-8h12a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1H2v6Zm7 3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H10a1 1 0 0 1-1-1v-4Z"/>`),
		g.Group(children),
	)
}