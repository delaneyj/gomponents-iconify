package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeglassOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.536 5.546L3 14M16 4h2l3 10m-11 2h4m5.426 3.423A3.5 3.5 0 0 1 14 16.5V14m4 0h3v2.5c0 .157-.01.312-.03.463M10 16.5a3.5 3.5 0 0 1-7 0V14h7v2.5M3 3l18 18"/>`),
		g.Group(children),
	)
}