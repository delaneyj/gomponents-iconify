package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastReverseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15 15.038V9.69c0-.882-1.058-1.332-1.693-.72l-6.558 6.31a1 1 0 0 0 0 1.44l6.558 6.31c.635.612 1.693.162 1.693-.72v-5.348l6.307 6.069c.635.611 1.693.16 1.693-.72V9.69c0-.882-1.058-1.332-1.693-.72L15 15.037Z"/><path d="M1 6a5 5 0 0 1 5-5h20a5 5 0 0 1 5 5v20a5 5 0 0 1-5 5H6a5 5 0 0 1-5-5V6Zm5-3a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Z"/></g>`),
		g.Group(children),
	)
}