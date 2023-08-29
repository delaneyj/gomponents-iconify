package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Element(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.717 11.431C11.208 11.431 3.5 19.14 3.5 28.65m23.783 7.973c9.509 0 17.218-7.709 17.218-17.218M11.42 27.282c0 9.51 7.708 17.218 17.218 17.218m7.973-23.782c0-9.51-7.708-17.218-17.217-17.218"/>`),
		g.Group(children),
	)
}