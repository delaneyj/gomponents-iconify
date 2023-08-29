package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shortwave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 18.652c18.6 32.575 26.213-1.869 39 10.94c-21.055-33.378-24.021 2.356-39-10.94Z"/>`),
		g.Group(children),
	)
}