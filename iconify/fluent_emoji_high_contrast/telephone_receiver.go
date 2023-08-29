package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelephoneReceiver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.72 25.84l-7.13-7.13a.967.967 0 0 0-1.18-.14a6.198 6.198 0 0 0-2.7 3.43c-.23.72-1.07 1.05-1.72.66l-.01-.01l-.01-.01a22.171 22.171 0 0 1-7.61-7.61c0-.01-.01-.02-.01-.02c-.4-.65-.07-1.49.65-1.72a6.13 6.13 0 0 0 3.43-2.7c.23-.38.17-.86-.14-1.18L6.16 2.28a.967.967 0 0 0-1.18-.14a6.127 6.127 0 0 0-2.97 5.01v.01c-.01.1-.01.2-.01.3c0 .14.01.27.02.4a23.16 23.16 0 0 0 3.29 11.22c1.87 3.12 4.49 5.74 7.61 7.61l.01.01c3.44 2.07 7.36 3.2 11.37 3.29c.08 0 .16.01.24.01h.56v-.02a6.187 6.187 0 0 0 4.77-2.96c.22-.38.16-.87-.15-1.18Z"/>`),
		g.Group(children),
	)
}