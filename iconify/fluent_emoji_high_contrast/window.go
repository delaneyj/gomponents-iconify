package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Window(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M6.5 5A1.5 1.5 0 0 0 5 6.5v8a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5h-8Zm11 0a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5v-8A1.5 1.5 0 0 0 25.5 5h-8Zm8.354 1.146A.5.5 0 0 1 26 6.5V14h-8l7.854-7.854ZM17 17.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v8a1.5 1.5 0 0 1-1.5 1.5h-8a.5.5 0 0 1-.5-.5v-9Zm1 8.5h7.5a.5.5 0 0 0 .5-.5V18h-8v8ZM5.5 17a.5.5 0 0 0-.5.5v8A1.5 1.5 0 0 0 6.5 27h8a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5h-9Zm8.5 1v8H6.5a.499.499 0 0 1-.354-.146L14 18Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}