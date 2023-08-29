package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fitx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.11 24.062l10.307 10.565c-3.146 2.755-7.49 4.829-12.494 5.83l-6.913-7.078l-6.923 7.088c-5.004-1.001-9.358-3.065-12.505-5.83L14.91 24.062L4.5 13.404c3.126-2.765 7.46-4.86 12.463-5.87l7.047 7.211l7.037-7.212c4.993 1.022 9.327 3.116 12.453 5.881L33.11 24.062Z"/>`),
		g.Group(children),
	)
}