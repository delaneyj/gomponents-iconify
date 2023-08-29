package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mango(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00D26A" d="M5.535 2.49L1 4.788C2.113 7.544 5.014 8.756 7.465 7.51L12.5 5c-1.113-2.756-4.75-3.65-6.965-2.51Z"/><path fill="#FFB02E" d="m11.343 4.798l-1.17.534C6.17 7.168 4.124 11.715 5.357 15.984a26.945 26.945 0 0 0 5.206 9.787l.885 1.068a8.738 8.738 0 0 0 10.38 2.359c3.794-1.74 5.86-5.935 4.964-10.055l-.348-1.58a27.015 27.015 0 0 0-3.53-8.645l-.495-.779a8.742 8.742 0 0 0-11.076-3.34Z"/></g>`),
		g.Group(children),
	)
}