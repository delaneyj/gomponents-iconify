package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M32 33H42C43.1046 33 44 32.1046 44 31V15C44 13.8954 43.1046 13 42 13H32"/><path d="M16 33H6C4.89543 33 4 32.1046 4 31V15C4 13.8954 4.89543 13 6 13H16"/><path d="M12 13V19H36V13"/><path d="M16 13V5H32V13"/><path d="M16 29V43H32V29"/><path d="M22 35H26"/><line x1="13" x2="35" y1="27" y2="27"/></g>`),
		g.Group(children),
	)
}