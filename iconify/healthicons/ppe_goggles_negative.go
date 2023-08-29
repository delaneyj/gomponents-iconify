package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeGogglesNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPpeGogglesNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM4 22C4 14.82 9.82 9 17 9h14c7.18 0 13 5.82 13 13a5.002 5.002 0 0 1-4 4.9V29a8 8 0 0 1-15.653 2.336a.558.558 0 0 0-.17-.264A.274.274 0 0 0 24 31c-.05 0-.11.017-.177.072a.558.558 0 0 0-.17.264A8 8 0 0 1 8 29v-2.1A5.002 5.002 0 0 1 4 22Zm38 0a3.001 3.001 0 0 1-2.128 2.871A5.002 5.002 0 0 0 35 21H13a5.002 5.002 0 0 0-4.872 3.871A3.001 3.001 0 0 1 6 22c0-6.075 4.925-11 11-11h14c6.075 0 11 4.925 11 11Zm-32 4a3 3 0 0 1 3-3h22a3 3 0 0 1 3 3v3a6 6 0 0 1-11.74 1.753C25.98 29.833 25.14 29 24 29c-1.141 0-1.98.834-2.26 1.753A6 6 0 0 1 10 29v-3Zm3-1.5a1.5 1.5 0 0 0-1.5 1.5v3a4.5 4.5 0 0 0 8.805 1.315c.43-1.412 1.76-2.815 3.695-2.815c1.935 0 3.264 1.403 3.695 2.815A4.5 4.5 0 0 0 36.5 29v-3a1.5 1.5 0 0 0-1.5-1.5H13Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPpeGogglesNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}