package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chestnut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#F3AD61" d="M29.68 22.172L16 17L2.32 22.172c1.04 3.97 4.52 7.28 9.39 7.28h8.57c4.88 0 8.36-3.31 9.4-7.28Z"/><path fill="#6D4534" d="M29.68 22.172c.83-3.16.11-6.75-2.69-9.42l-9.59-9.19c-.78-.75-2.02-.75-2.8 0l-9.59 9.19c-2.8 2.68-3.51 6.26-2.69 9.42h27.36Z"/></g>`),
		g.Group(children),
	)
}