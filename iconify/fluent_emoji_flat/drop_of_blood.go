package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropOfBlood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#F8312F" d="M25.5 22.962c.9-2.88.4-5.52-.8-7.72l-6.93-12.2c-.79-1.39-2.79-1.39-3.58 0l-6.84 12.03c-.83 1.46-1.35 3.13-1.35 4.94c0 6.09 5.45 10.91 11.74 9.83c3.64-.62 6.66-3.36 7.76-6.88Z"/>`),
		g.Group(children),
	)
}