package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lipstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#F70A8D" d="M22.15 20.384H11.14l-.02-16.03c0-1.6 1.02-2.46 2.17-2.46c4.25 0 8.86 6.05 8.86 9.44v9.05Z"/><path fill="#FF6DC6" d="M21.503 12.955c1.462-1.34.488-4.782-2.175-7.69c-2.664-2.908-6.009-4.18-7.47-2.84c-1.463 1.34-.489 4.782 2.175 7.69c2.664 2.908 6.008 4.18 7.47 2.84Z"/><path fill="#FFB02E" d="M9 24.16v-3.08c0-.59.48-1.08 1.08-1.08h12.84c.59 0 1.08.48 1.08 1.08v3.08H9Z"/><path fill="#321B41" d="M7 30v-4.52c0-.81.67-1.48 1.48-1.48h16.04c.81 0 1.48.67 1.48 1.48V30H7Z"/></g>`),
		g.Group(children),
	)
}