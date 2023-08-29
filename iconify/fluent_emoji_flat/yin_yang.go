package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#8D65C5" d="M30 26a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20Z"/><path fill="#fff" d="M16 24.5h-.031a4.25 4.25 0 0 0 0-8.5a4.25 4.25 0 1 1-.076-8.5H16a8.5 8.5 0 0 1 0 17Zm0 1.5a10 10 0 1 0 0-20a10 10 0 0 0 0 20Zm-.031-3.625a2.125 2.125 0 1 1 0-4.25a2.125 2.125 0 0 1 0 4.25ZM16 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`),
		g.Group(children),
	)
}