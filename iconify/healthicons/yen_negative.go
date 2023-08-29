package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YenNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsYenNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM35.536 13.28a2 2 0 1 0-3.072-2.56L24 20.876L15.537 10.72a2 2 0 1 0-3.073 2.56L19.73 22H18a2 2 0 0 0 0 4h4v2h-4a2 2 0 0 0 0 4h4v4a2 2 0 0 0 4 0v-4h4a2 2 0 0 0 0-4h-4v-2h4a2 2 0 0 0 0-4h-1.73l7.266-8.72Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsYenNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}