package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m17.678 2.829l11.55 12.25c1.27 1.35.79 3.56-.92 4.25l-4.213 1.707l.97 2.369v.002a2.607 2.607 0 0 1-1.428 3.42l-4.895 2.008h-.001a2.607 2.607 0 0 1-3.42-1.428l-.99-2.416l-4.563 1.848c.62.65.53 1.62-.13 2.23c-.65.61-1.68.58-2.29-.07l-4.91-5.21c-.61-.65-.58-1.68.07-2.29c.65-.61 1.68-.58 2.29.07l8.59-18.06c.79-1.68 3.02-2.03 4.29-.68Zm-1.63 21.466l.99 2.418a.755.755 0 0 0 .996.41l.003-.001l4.906-2.012a.755.755 0 0 0 .41-.997l-.002-.003l-.973-2.379l-6.33 2.564Z"/>`),
		g.Group(children),
	)
}