package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FemaleSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M16 5a7.67 7.67 0 0 0-1 15.275V21.7h-2.33a1 1 0 1 0 0 2H15V26a1 1 0 1 0 2 0v-2.3h2.33a1 1 0 0 0 0-2H17v-1.425A7.671 7.671 0 0 0 16 5Zm-5.67 7.67a5.67 5.67 0 1 1 11.34 0a5.67 5.67 0 0 1-11.34 0Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}