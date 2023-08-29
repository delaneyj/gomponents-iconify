package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparkle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00D26A" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="m17.209 8.008l-.704-2.392a.5.5 0 0 0-.96 0l-.704 2.392a10 10 0 0 1-6.769 6.77l-2.392.704c-.479.14-.479.818 0 .959l2.392.704a10 10 0 0 1 6.77 6.77l.703 2.392c.141.478.819.478.96 0l.704-2.393a10 10 0 0 1 6.77-6.769l2.392-.704c.478-.14.478-.818 0-.96l-2.393-.704a10 10 0 0 1-6.769-6.769Z"/></g>`),
		g.Group(children),
	)
}