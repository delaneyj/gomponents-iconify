package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.21 14.339L-.007 8.22l3.084-3.035L6.21 8.268l6.713-6.607l3.084 3.035l-9.797 9.643zM1.686 8.22l4.524 4.453l8.104-7.976l-1.391-1.369L6.21 9.935L3.077 6.852L1.686 8.221z"/>`),
		g.Group(children),
	)
}