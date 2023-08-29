package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboardcleaner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 4.5L30.91 17.09m0 0c3.03 3.03 4.213 6.91 2.356 11.198c-1.747 4.034-6.52 15.212-6.52 15.212s-7.708-.515-14.72-7.526S4.5 21.254 4.5 21.254s11.178-4.773 15.212-6.52C24 12.877 27.881 14.06 30.91 17.09Zm-12.751-1.685L32.504 29.75"/>`),
		g.Group(children),
	)
}