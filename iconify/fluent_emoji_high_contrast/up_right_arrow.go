package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M22.01 9a1 1 0 0 1 1.006 1.005l-.033 6.441c-.005.889-1.079 1.33-1.707.702l-1.493-1.493c-.127-.127-.349-.08-.472.05a.777.777 0 0 1-.018.018l-7.586 7.585a1 1 0 0 1-1.414 0l-1.586-1.585a1 1 0 0 1 0-1.415l7.586-7.585l.018-.018c.13-.124.176-.345.05-.472l-1.494-1.493c-.628-.628-.186-1.703.702-1.707L22.011 9Z"/><path fill-rule="evenodd" d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}