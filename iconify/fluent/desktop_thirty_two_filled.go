package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.25A3.25 3.25 0 0 1 5.25 3h21.5A3.25 3.25 0 0 1 30 6.25v15.5A3.25 3.25 0 0 1 26.75 25h-6.744v2.001h2.998a1 1 0 1 1 0 1.999H9.012a1 1 0 1 1 0-1.999h2.998V25H5.25A3.25 3.25 0 0 1 2 21.75V6.25ZM14.01 25v2.001h3.996V25H14.01Z"/>`),
		g.Group(children),
	)
}