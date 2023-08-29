package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DenoSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 7H6V6h1v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM8 13.981a6.462 6.462 0 0 0 2.971-.985l-.287-5.167A2.995 2.995 0 0 0 7.694 5H6a2 2 0 0 0-1.732 1H5v1H4a2 2 0 0 0 2 2h.882c.496 0 .95-.28 1.17-.724l.895.448A2.308 2.308 0 0 1 8 9.71v4.27Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}