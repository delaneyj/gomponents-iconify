package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardClockThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 5A4.5 4.5 0 0 0 2 9.5V11h28V9.5A4.5 4.5 0 0 0 25.5 5h-19ZM2 22.5V13h28v4.343A9 9 0 0 0 14.935 27H6.5A4.5 4.5 0 0 1 2 22.5Zm21 8a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15ZM22.75 18a.75.75 0 0 0-.75.75v4.5c0 .414.336.75.75.75h3.5a.75.75 0 0 0 0-1.5H23.5v-3.75a.75.75 0 0 0-.75-.75Z"/>`),
		g.Group(children),
	)
}