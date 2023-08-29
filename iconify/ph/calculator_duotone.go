package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M176 64v48H80V64Z" opacity=".2"/><path d="M80 120h96a8 8 0 0 0 8-8V64a8 8 0 0 0-8-8H80a8 8 0 0 0-8 8v48a8 8 0 0 0 8 8Zm8-48h80v32H88Zm112-48H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16Zm0 192H56V40h144Zm-100-68a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm40 0a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm40 0a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm-80 40a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm40 0a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm40 0a12 12 0 1 1-12-12a12 12 0 0 1 12 12Z"/></g>`),
		g.Group(children),
	)
}