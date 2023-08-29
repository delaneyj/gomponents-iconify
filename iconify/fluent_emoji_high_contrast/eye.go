package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M21 16a5 5 0 1 1-9.643-1.86a2 2 0 1 0 2.784-2.784A5 5 0 0 1 21 16Z"/><path d="M15.984 24.969a8.984 8.984 0 1 0 0-17.97a8.984 8.984 0 0 0 0 17.97Zm0-2a6.984 6.984 0 1 1 0-13.97a6.984 6.984 0 0 1 0 13.97Z"/><path d="M16.156 30.313c7.819 0 14.157-6.338 14.157-14.157C30.313 8.338 23.974 2 16.155 2C8.338 2 2 8.338 2 16.156c0 7.819 6.338 14.157 14.156 14.157Zm0-2C9.443 28.313 4 22.87 4 16.155C4 9.443 9.443 4 16.156 4c6.714 0 12.157 5.443 12.157 12.156c0 6.714-5.443 12.157-12.157 12.157Z"/></g>`),
		g.Group(children),
	)
}