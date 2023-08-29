package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 24L41 24"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 7L24 41"/><rect width="4" height="4" x="5" y="5" fill="#000"/><rect width="4" height="4" x="14" y="5" fill="#000"/><rect width="4" height="4" x="30" y="5" fill="#000"/><rect width="4" height="4" x="39" y="5" fill="#000"/><rect width="4" height="4" x="39" y="14" fill="#000"/><rect width="4" height="4" x="5" y="14" fill="#000"/><rect width="4" height="4" x="5" y="39" fill="#000"/><rect width="4" height="4" x="14" y="39" fill="#000"/><rect width="4" height="4" x="30" y="39" fill="#000"/><rect width="4" height="4" x="39" y="39" fill="#000"/><rect width="4" height="4" x="39" y="30" fill="#000"/><rect width="4" height="4" x="5" y="30" fill="#000"/></g>`),
		g.Group(children),
	)
}