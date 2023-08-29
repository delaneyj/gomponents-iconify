package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiEditSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a5.997 5.997 0 0 1 5.198 3a2.862 2.862 0 0 0-2.096.84l-.734.735a.75.75 0 1 0-1.043 1.043l-2.558 2.557a2.512 2.512 0 0 1-.63-.508a.5.5 0 1 0-.746.667c.196.218.42.412.664.576c-.296.364-.51.788-.623 1.245l-.282 1.126A6 6 0 0 1 8 2ZM6.25 7.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm.73 3.627l4.83-4.83a1.87 1.87 0 1 1 2.644 2.646l-4.83 4.829a2.197 2.197 0 0 1-1.02.578l-1.498.374a.89.89 0 0 1-1.079-1.078l.375-1.498a2.18 2.18 0 0 1 .578-1.02Z"/>`),
		g.Group(children),
	)
}