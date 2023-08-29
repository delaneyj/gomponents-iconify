package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 29h7.5a4.5 4.5 0 0 0 4.5-4.5V17H17v12Zm0-14h12V7.5A4.5 4.5 0 0 0 24.5 3H17v12ZM15 3v12H3V7.5A4.5 4.5 0 0 1 7.5 3H15Zm0 14v12H7.5A4.5 4.5 0 0 1 3 24.5V17h12Z"/>`),
		g.Group(children),
	)
}