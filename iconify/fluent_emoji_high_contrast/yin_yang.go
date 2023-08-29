package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M18 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path d="M26 16c0 5.523-4.477 10-10 10S6 21.523 6 16S10.477 6 16 6s10 4.477 10 10Zm-1.5 0a8.5 8.5 0 0 0-8.607-8.5a4.25 4.25 0 0 0 .076 8.5a4.25 4.25 0 1 1 0 8.5H16a8.5 8.5 0 0 0 8.5-8.5Zm-6.406 4.25a2.125 2.125 0 1 0-4.25 0a2.125 2.125 0 0 0 4.25 0Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}