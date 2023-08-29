package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TumblerGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M5 25.444v-6.456c0-.549.39-.988.878-.988H7v-4.007c0-.55.214-.993.48-.993h7.04c.266 0 .48.443.48.993V18h1v-4.002c0-.553.428-.998.96-.998h6.08c.532 0 .96.445.96.998V18h1V7a1 1 0 1 1 2 0v11h.122c.488 0 .878.439.878.988v6.456c0 .858-.62 1.556-1.383 1.556H6.383C5.621 27 5 26.302 5 25.444ZM25 18v7a1 1 0 1 0 2 0v-7h-2Z"/><path d="M1 4.621C1 3.728 1.729 3 2.621 3H29.38c.89 0 1.62.728 1.62 1.621v22.874A3.497 3.497 0 0 1 27.505 31h-23A3.506 3.506 0 0 1 1 27.495V4.621ZM3 5v22.495C3 28.325 3.674 29 4.505 29h23C28.33 29 29 28.33 29 27.495V5H3Z"/></g>`),
		g.Group(children),
	)
}