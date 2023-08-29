package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#FFB02E" d="m25.064 23.405l-1.148-2.807l-1.715.702l1.15 2.81l.002.003a.755.755 0 0 1-.41.997l-4.906 2.012l-.003.001a.755.755 0 0 1-.996-.41l-1.152-2.815l-1.715.702l1.15 2.807a2.607 2.607 0 0 0 3.42 1.428h.002l4.894-2.008a2.607 2.607 0 0 0 1.428-3.42v-.002ZM8 23.5l1.708 3.279c.62.65.59 1.68-.07 2.29c-.65.61-1.68.58-2.29-.07l-4.91-5.21c-.61-.65-.58-1.68.07-2.29c.65-.61 1.68-.58 2.29.07L8 23.5Z"/><path fill="#F9C23C" d="m29.228 15.079l-11.55-12.25c-1.27-1.35-3.5-1-4.29.68l-8.59 18.06l4.97 5.27l18.54-7.51c1.71-.69 2.19-2.9.92-4.25Z"/></g>`),
		g.Group(children),
	)
}