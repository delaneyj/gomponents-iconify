package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparkle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="m16.505 5.616l.704 2.392a10 10 0 0 0 6.77 6.77l2.392.704c.478.14.478.818 0 .959l-2.393.704a10 10 0 0 0-6.769 6.77l-.704 2.392c-.141.478-.819.478-.96 0l-.704-2.393a10 10 0 0 0-6.769-6.769l-2.392-.704c-.479-.14-.479-.818 0-.96l2.392-.704a10 10 0 0 0 6.77-6.769l.703-2.392a.5.5 0 0 1 .96 0Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}