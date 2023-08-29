package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 42H40a14 14 0 0 0-14 14v144a14 14 0 0 0 14 14h176a14 14 0 0 0 14-14V56a14 14 0 0 0-14-14ZM40 54h176a2 2 0 0 1 2 2v42H38V56a2 2 0 0 1 2-2Zm-2 146v-90h60v92H40a2 2 0 0 1-2-2Zm178 2H110v-92h108v90a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}