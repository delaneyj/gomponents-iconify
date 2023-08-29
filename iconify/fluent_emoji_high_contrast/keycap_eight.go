package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M11.456 12.484a4.516 4.516 0 1 1 8.073 2.783a5.237 5.237 0 1 1-7.091.03a4.497 4.497 0 0 1-.982-2.813Zm4.516-1.266a1.266 1.266 0 1 0 0 2.533a1.266 1.266 0 0 0 0-2.533ZM16 17.399a1.737 1.737 0 1 0 0 3.474a1.737 1.737 0 0 0 0-3.474Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}