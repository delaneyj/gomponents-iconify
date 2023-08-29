package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeSuitNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPpeSuitNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM22 4a4 4 0 0 0-4 4v6.417h-1a4 4 0 0 0-4 4V31a3 3 0 0 0 5 2.236V40.5a2.5 2.5 0 0 0 5 0V29h2v11.5a2.5 2.5 0 0 0 5 0v-7.264A3 3 0 0 0 35 31V18.417a4 4 0 0 0-4-4h-1V8a4 4 0 0 0-4-4h-4Zm.9 19.792h-5.5v2.083h13.2v-2.083h-5.5v-9.375h-2.2v9.375ZM20 7h8v2a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V7Zm-5 23v1a1 1 0 1 0 2 0v-1h-2Zm16 1v-1h2v1a1 1 0 1 1-2 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPpeSuitNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}