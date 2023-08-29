package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Easypark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.838 29.377l20.544-4.943c10.757-2.588 9.506-15.966 1.453-18.315C20.35 2.769 9.838 4.485 9.838 12.515v23.451c0 13.2 26.746 5.04 26.746 5.04"/>`),
		g.Group(children),
	)
}