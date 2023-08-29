package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M6.063 14.627a5.268 5.268 0 0 1 9.766-2.74l.033.054a1.375 1.375 0 1 1-2.349 1.43l-.033-.053a2.517 2.517 0 0 0-4.668 1.309v2.76a2.536 2.536 0 0 0 4.687 1.342l.022-.035a1.375 1.375 0 1 1 2.333 1.456l-.022.035a5.286 5.286 0 0 1-9.77-2.799v-2.76ZM19.375 9.36c.76 0 1.375.615 1.375 1.374v8.907c0 .138.112.25.25.25h3.656a1.375 1.375 0 1 1 0 2.75h-5.281c-.76 0-1.375-.616-1.375-1.375V10.734c0-.759.616-1.375 1.375-1.375Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}