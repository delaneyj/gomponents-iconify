package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagClockThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 3A1.5 1.5 0 0 0 5 4.5V28a1 1 0 1 0 2 0v-7h7.223a9.003 9.003 0 0 1 10.678-6.799L23.25 12l5.55-7.4A1 1 0 0 0 28 3H6.5ZM23 30.5a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15ZM22.75 18a.75.75 0 0 0-.75.75v4.5c0 .414.336.75.75.75h3.5a.75.75 0 0 0 0-1.5H23.5v-3.75a.75.75 0 0 0-.75-.75Z"/>`),
		g.Group(children),
	)
}