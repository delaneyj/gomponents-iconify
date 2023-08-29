package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#F3AD61" d="M16.331 2c-2.57 0-4.98 1.28-6.42 3.4l-.13.19A21.923 21.923 0 0 0 6.101 20l.08.8c.5 5.21 4.88 9.18 10.11 9.18c5.25 0 9.63-3.99 10.11-9.22l.09-.93a21.9 21.9 0 0 0-3.78-14.48A7.772 7.772 0 0 0 16.331 2Z"/>`),
		g.Group(children),
	)
}