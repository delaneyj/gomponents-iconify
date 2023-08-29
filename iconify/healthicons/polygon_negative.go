package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolygonNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPolygonNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm17.874 39A4.002 4.002 0 0 1 10 38a4 4 0 0 1 2.184-3.565l-2.32-10.437a4 4 0 1 1 2.706-7.063l7.696-5.497a4 4 0 1 1 7.467 0l7.697 5.497a4 4 0 1 1 2.706 7.063l-2.32 10.437A4 4 0 1 1 30.126 39H17.874Zm0-2a4.002 4.002 0 0 0-3.738-2.998l-2.32-10.437a4 4 0 0 0 1.918-5.003l7.696-5.497A3.984 3.984 0 0 0 24 14a3.98 3.98 0 0 0 2.57-.935l7.696 5.497a4 4 0 0 0 1.917 5.003l-2.319 10.437A4.002 4.002 0 0 0 30.126 37H17.874Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPolygonNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}