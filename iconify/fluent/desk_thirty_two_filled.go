package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 4A3.25 3.25 0 0 0 2 7.25V10h26v17a1 1 0 1 0 2 0V7.25A3.25 3.25 0 0 0 26.75 4H5.25ZM16 12H2v11.5A4.5 4.5 0 0 0 6.5 28h5a4.5 4.5 0 0 0 4.5-4.5V12ZM6 17a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Z"/>`),
		g.Group(children),
	)
}