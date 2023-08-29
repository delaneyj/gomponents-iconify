package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mindfulnotifier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="20.721" cy="17.963" r="7.611" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.721 29.76c8.467 0 15.221 2.367 15.221 5.19v7.264H5.5V34.95c0-2.823 6.753-5.19 15.221-5.19ZM32.974 9.702a14.787 14.787 0 0 1 0 16.523M38.78 5.786a21.797 21.797 0 0 1 0 24.355"/>`),
		g.Group(children),
	)
}