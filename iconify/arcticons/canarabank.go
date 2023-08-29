package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Canarabank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.935 31.709l-5.675-.7L30.2 8.724l13.3 25.52l-15.178-1.871"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.211 24.763L4.5 27.055l15.366 12.221l13.12-16.983l-8.939 1.494"/>`),
		g.Group(children),
	)
}