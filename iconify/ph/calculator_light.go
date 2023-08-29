package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 58H80a6 6 0 0 0-6 6v48a6 6 0 0 0 6 6h96a6 6 0 0 0 6-6V64a6 6 0 0 0-6-6Zm-6 48H86V70h84Zm30-80H56a14 14 0 0 0-14 14v176a14 14 0 0 0 14 14h144a14 14 0 0 0 14-14V40a14 14 0 0 0-14-14Zm2 190a2 2 0 0 1-2 2H56a2 2 0 0 1-2-2V40a2 2 0 0 1 2-2h144a2 2 0 0 1 2 2ZM98 148a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm40 0a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm40 0a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm-80 40a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm40 0a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm40 0a10 10 0 1 1-10-10a10 10 0 0 1 10 10Z"/>`),
		g.Group(children),
	)
}