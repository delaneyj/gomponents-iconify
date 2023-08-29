package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scorpio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="m25.55 16.55l2 2.12c.28.31.27.78-.04 1.07a.751.751 0 0 1-1.06-.04l-.45-.48v2.28c0 1.93-1.57 3.5-3.5 3.5S19 23.43 19 21.5V11c0-1.1-.9-2-2-2s-2 .9-2 2v13c0 .55-.45 1-1 1s-1-.45-1-1V11c0-1.1-.9-2-2-2s-2 .9-2 2v13c0 .55-.45 1-1 1s-1-.45-1-1V11c0-1.1-.9-2-2-2c-.55 0-1-.45-1-1s.45-1 1-1a4 4 0 0 1 3 1.36A4 4 0 0 1 11 7a4 4 0 0 1 3 1.36A4 4 0 0 1 17 7c2.21 0 4 1.79 4 4v10.5a1.5 1.5 0 0 0 3 .02v-2.3l-.46.48c-.28.3-.76.31-1.06.03a.755.755 0 0 1-.03-1.06l2-2.12c.15-.15.34-.24.55-.24c.21 0 .41.09.55.24Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}