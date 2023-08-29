package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBug0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBug1" fill="currentColor"><path id="feBug2" d="M6 15H4a1 1 0 0 1 0-2h2v-1.337a1 1 0 0 1-1-.249L3.586 10A1 1 0 0 1 5 8.586l1.136 1.136a6.007 6.007 0 0 1 2.665-3.8L7.586 4.708A1 1 0 0 1 9 3.293l1.823 1.822a6.029 6.029 0 0 1 2.354 0L15 3.293a1 1 0 0 1 1.414 1.414L15.2 5.923a6.007 6.007 0 0 1 2.665 3.8L19 8.585A1 1 0 0 1 20.414 10L19 11.414a1 1 0 0 1-1 .25V13h2a1 1 0 0 1 0 2h-2c0 .483-.057.953-.165 1.404a1 1 0 0 1 1.165.182L20.414 18A1 1 0 0 1 19 19.414L17.586 18a.997.997 0 0 1-.216-.321A6 6 0 0 1 12 21a6 6 0 0 1-5.37-3.321a.997.997 0 0 1-.216.321L5 19.414A1 1 0 1 1 3.586 18L5 16.586a1 1 0 0 1 1.165-.182A6.016 6.016 0 0 1 6 15Zm9.874-5H8.126a4.002 4.002 0 0 1 7.748 0Zm0 6a4.002 4.002 0 0 1-7.748 0h7.748Z"/></g></g>`),
		g.Group(children),
	)
}