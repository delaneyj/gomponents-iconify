package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMinusSquareFill0"><g id="evaMinusSquareFill1"><path id="evaMinusSquareFill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm-3 10H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2Z"/></g></g>`),
		g.Group(children),
	)
}