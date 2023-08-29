package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OktaVerify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.013 12.257a21.53 21.53 0 1 1-1.676-2.234"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.699 19.775a11.513 11.513 0 1 1-1.473-2.641"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.336 10.024L24 26.36l-4.72-4.72"/>`),
		g.Group(children),
	)
}