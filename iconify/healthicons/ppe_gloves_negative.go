package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeGlovesNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPpeGlovesNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM25.5 5A1.5 1.5 0 0 0 24 6.5V19h-1V8.5a1.5 1.5 0 0 0-3 0V19h-1v-8.5a1.5 1.5 0 0 0-3 0V27a5.997 5.997 0 0 0 2.905 5.141L18 43h12l-.91-10.91a9.116 9.116 0 0 0 3.617-3.15l4.25-6.375a1.881 1.881 0 0 0-3.034-2.218L31 24V8.5a1.5 1.5 0 0 0-3 0V19h-1V6.5A1.5 1.5 0 0 0 25.5 5ZM14 12h-2v16c0 2.286 1.18 4.647 2.943 5.989L14.205 41h2.011l.845-8.028l-.544-.328C15.092 31.784 14 29.85 14 28V12Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPpeGlovesNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}